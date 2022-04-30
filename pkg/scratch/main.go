package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scratch/pretender"
	"k8s.io/kubernetes/pkg/scratch/pretender/podtraits"
)

func InitLogs() {
	klog.InitFlags(nil)
	flag.Parse()
}

func NewAntiAffinityPod(name string) pretender.PodWithTraits {
	return pretender.PodWithTraits{
		Pod: &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				UID:       types.UID(name),
				Namespace: "global-namespace",
				Labels: map[string]string{
					"name":                name,
					"anti-affinity-group": "1",
				},
			},
			Spec: v1.PodSpec{
				Affinity: &v1.Affinity{
					PodAntiAffinity: &v1.PodAntiAffinity{
						PreferredDuringSchedulingIgnoredDuringExecution: []v1.WeightedPodAffinityTerm{
							{
								Weight: 100,
								PodAffinityTerm: v1.PodAffinityTerm{
									LabelSelector: &metav1.LabelSelector{
										MatchLabels:      map[string]string{"anti-affinity-group": "1"},
										MatchExpressions: []metav1.LabelSelectorRequirement{},
									},
									Namespaces:        []string{"global-namespace"},
									TopologyKey:       "name",
									NamespaceSelector: nil,
								},
							},
						},
					},
				},
			},
		},
		Traits: []pretender.PodTrait{
			podtraits.AffectNodeCount{},
		},
	}
}

func NewResourceRequestingPod(name string) pretender.PodWithTraits {
	return pretender.PodWithTraits{
		Pod: &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				UID:       types.UID(name),
				Namespace: "global-namespace",
				Labels: map[string]string{
					"name": name,
				},
			},
		},
		Traits: []pretender.PodTrait{
			podtraits.AffectNodeCount{},
			podtraits.RequestMemory{Request: 1024},
			podtraits.RequestCPU{Request: 4096},
		},
	}
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

func addNode(ctx context.Context, fwk framework.Framework, sched *scheduler.Scheduler, node *v1.Node) {
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
	ctx := context.Background()
	snapshot := scheduler.NewSnapshot()
	sched := scheduler.CreateTestScheduler(ctx, snapshot)
	ps := pretender.NewStateManager(sched)
	fwk := scheduler.NewTestFramework(pretender.NewClientset(&ps), snapshot)
	err := ps.SetFramework(fwk)
	if err != nil {
		panic(err)
	}
	var ret []pretender.StateSnapshot

	addNode(ctx, fwk, sched, NewTestNode("My node #1"))
	addNode(ctx, fwk, sched, NewTestNode("My node #2"))
	addNode(ctx, fwk, sched, NewTestNode("My node #3"))

	// schedule some pods
	var pods []pretender.PodWithTraits
	for i := 0; i < 16; i++ {
		pods = append(pods, NewResourceRequestingPod(fmt.Sprintf("pod%d", i)))
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
