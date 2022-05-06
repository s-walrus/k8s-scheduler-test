package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/defaultbinder"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/feature"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/interpodaffinity"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/noderesources"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/queuesort"
	st "k8s.io/kubernetes/pkg/scheduler/testing"
	"k8s.io/kubernetes/pkg/scratch/pretender"
	"k8s.io/kubernetes/pkg/scratch/pretender/podbuilder"
)

func InitLogs() {
	klog.InitFlags(nil)
	flag.Parse()
}

func NewTestNode(name string) *v1.Node {
	node := v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			UID:  types.UID("my node"),
			Labels: map[string]string{
				"name": name,
			},
		},
		Status: v1.NodeStatus{
			Capacity: v1.ResourceList{
				v1.ResourceCPU:    *(resource.NewQuantity(100500, resource.DecimalSI)),
				v1.ResourceMemory: *(resource.NewQuantity(100500, resource.DecimalSI)),
				v1.ResourcePods:   *(resource.NewQuantity(10, resource.DecimalSI)),
			},
			Allocatable: v1.ResourceList{
				v1.ResourceCPU:    *(resource.NewQuantity(100500, resource.DecimalSI)),
				v1.ResourceMemory: *(resource.NewQuantity(100500, resource.DecimalSI)),
				v1.ResourcePods:   *(resource.NewQuantity(10, resource.DecimalSI)),
			}},
	}
	return &node
}

func NewInterPodAffinity(plArgs runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return interpodaffinity.New(plArgs, h, feature.Features{
		EnablePodAffinityNamespaceSelector: true,
		EnablePodDisruptionBudget:          false,
		EnablePodOverhead:                  false,
		EnableReadWriteOncePod:             false,
		EnableVolumeCapacityPriority:       false,
		EnableCSIStorageCapacity:           false,
	})
}

func NewBalancedAllocation(plArgs runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return noderesources.NewBalancedAllocation(plArgs, h, feature.Features{
		EnablePodAffinityNamespaceSelector: true,
		EnablePodDisruptionBudget:          false,
		EnablePodOverhead:                  false,
		EnableReadWriteOncePod:             false,
		EnableVolumeCapacityPriority:       false,
		EnableCSIStorageCapacity:           false,
	})
}

func addNode(ctx context.Context, fwk framework.Framework, node *v1.Node) {
	_, err := fwk.ClientSet().CoreV1().Nodes().Create(ctx, node, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
}

func SchedulePodWithTraits(sched *scheduler.Scheduler, fwk framework.Framework, ps *pretender.StateManager, pod pretender.PodWithTraits) {
	ps.AddOrUpdatePod(pod)

	ctx := context.Background()
	scheduler.FakeScheduleOne(ctx, sched, fwk, pod.Pod)
}

func EvalSchedulerDemo() []pretender.StateSnapshot {
	plugins := []st.RegisterPluginFunc{
		st.RegisterQueueSortPlugin(queuesort.Name, queuesort.New),
		st.RegisterPluginAsExtensions(interpodaffinity.Name, NewInterPodAffinity, "PreFilter", "Filter", "PreScore", "Score"),
		st.RegisterPluginAsExtensions(noderesources.BalancedAllocationName, NewBalancedAllocation, "Score"),
		st.RegisterFilterPlugin("TrueFilter", st.NewTrueFilterPlugin),
		st.RegisterBindPlugin(defaultbinder.Name, defaultbinder.New),
	}
	ctx := context.Background()
	snapshot := scheduler.NewSnapshot()
	sched := scheduler.CreateTestScheduler(ctx, snapshot)
	ps := pretender.NewStateManager(sched)
	fwk := scheduler.NewTestFramework(pretender.NewClientset(&ps), snapshot, plugins)
	err := ps.SetFramework(fwk)
	if err != nil {
		panic(err)
	}
	var ret []pretender.StateSnapshot

	addNode(ctx, fwk, NewTestNode("My node #1"))
	addNode(ctx, fwk, NewTestNode("My node #2"))
	addNode(ctx, fwk, NewTestNode("My node #3"))

	selfAntiAffinityPodBuilder := podbuilder.NewPodBuilder("balanced")
	selfAntiAffinityPodBuilder.SetLabel("anti-affinity-group", "1")
	selfAntiAffinityPodBuilder.AddPreferredPodAntiAffinity(map[string]string{"anti-affinity-group": "1"})

	// schedule some pods
	var pods []pretender.PodWithTraits
	for i := 0; i < 16; i++ {
		pods = append(pods, selfAntiAffinityPodBuilder.GetPod())
	}
	for _, pod := range pods {
		SchedulePodWithTraits(sched, fwk, &ps, pod)
		ret = append(ret, ps.GetSnapshot())
	}

	return ret
}

func PrintTestResult(snapshots []pretender.StateSnapshot) {
	for _, s := range snapshots {
		fmt.Println("{")
		for node, state := range s {
			fmt.Printf(
				"\t%s: { cnt: %d, mem: %d, cpu: %d }\n",
				node,
				state.NodeCount,
				state.MemoryRequested,
				state.MilliCPURequested,
			)
		}
		fmt.Println("},")
	}
}

func main() {
	InitLogs()

	PrintTestResult(EvalSchedulerDemo())
}
