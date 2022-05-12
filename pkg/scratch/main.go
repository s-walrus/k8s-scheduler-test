package main

import (
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/defaultbinder"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/feature"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/interpodaffinity"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/noderesources"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/queuesort"
	st "k8s.io/kubernetes/pkg/scheduler/testing"
	"k8s.io/kubernetes/pkg/scratch/execution"
	"k8s.io/kubernetes/pkg/scratch/execution/requests"
	"k8s.io/kubernetes/pkg/scratch/pretender"
	"k8s.io/kubernetes/pkg/scratch/pretender/podbuilder"
)

func InitLogs() {
	klog.InitFlags(nil)
	flag.Parse()
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

func PrintTestResult(snapshots []pretender.StateSnapshot) {
	for _, s := range snapshots {
		fmt.Println("{")
		for node, state := range s {
			fmt.Printf(
				"\t%s: { cnt: %d, mem: %d, cpu: %d, t: %d }\n",
				node,
				state.NodeCount,
				state.MemoryRequested,
				state.MilliCPURequested,
				state.Time,
			)
		}
		fmt.Println("},")
	}
}

// TODO add node builder, remove NewTestNode

func NewTestNode(name string) *v1.Node {
	node := v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			UID:  uuid.NewUUID(),
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

func SelfAntiAffinityPodsScenario() *execution.StaticRequestGenerator {
	var reqs []execution.Request
	var time int64 = 0
	for i := 0; i < 3; i++ {
		reqs = append(reqs, requests.NewAddNode(NewTestNode(fmt.Sprintf("node%d", i+1)), time))
		time++
	}

	affinityPodBuilder := podbuilder.NewPodBuilder("affinity-pod")
	affinityPodBuilder.AddPreferredPodAntiAffinity(map[string]string{"affinity-group": "1"})
	affinityPodBuilder.SetLabel("affinity-group", "1")
	reqs = append(reqs, requests.NewRemoveNode("node2", time))
	time++
	var podUIDs []types.UID
	for i := 0; i < 4; i++ {
		pod := affinityPodBuilder.GetPod()
		podUIDs = append(podUIDs, pod.Pod.UID)
		reqs = append(reqs, requests.NewSchedulePod(pod, time))
		time++
	}
	for _, uid := range podUIDs {
		reqs = append(reqs, requests.NewKillPod(uid, time))
		time++
	}
	return execution.NewStaticRequestGenerator(reqs)
}

func main() {
	InitLogs()

	plugins := []execution.PluginInfo{
		execution.NewPluginInfo(queuesort.Name, queuesort.New, "QueueSort"),
		execution.NewPluginInfo(interpodaffinity.Name, NewInterPodAffinity, "PreFilter", "Filter", "PreScore", "Score"),
		execution.NewPluginInfo(noderesources.BalancedAllocationName, NewBalancedAllocation, "Score"),
		execution.NewPluginInfo("TrueFilter", st.NewTrueFilterPlugin, "Filter"),
		execution.NewPluginInfo(defaultbinder.Name, defaultbinder.New, "Bind"),
	}

	PrintTestResult(execution.RunSchedulerIsolationTest(plugins, SelfAntiAffinityPodsScenario()))
}

/*

Closest project goals:
+ test if updating pod resource requests is beneficial
+ [come up with a plugin and test it on a relatively large scale]

Tasks:
- add a pod trait that reflects real resource usage
	$ figure out how to do it with basic fourier series
	* add time consideration to pod traits handling
	- implement the pod trait
	* track choke time
+ create a scenario with pods consuming random real resources
	- implement a pod builder for random resource consuming pods (random are fourier series coefficients [realistic though])
	* add an option to make synchronized spikes in resource usage
	- add "request update" request
	- create the scenario and a similar one with request updates
* run the test with different configurations of the scenario (similar function shifted, for example)
+ make some useful metrics from test results
+ come up with a plugin and run the implemented scenario with it

*/
