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

func NewTestPod(name string) *v1.Pod {
	return &v1.Pod{
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
					RequiredDuringSchedulingIgnoredDuringExecution: []v1.PodAffinityTerm{
						{
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
	} else {
		sched.SchedulerCache.AddNode(node)
	}
}

func EvalSchedulerDemo() []pretender.StateSnapshot {
	ctx := context.Background()
	ps := pretender.NewState()
	snapshot := scheduler.NewSnapshot()
	sched := scheduler.CreateTestScheduler(ctx, snapshot)
	fwk := scheduler.NewTestFramework(&ps, snapshot)

	addNode(ctx, fwk, sched, NewTestNode("My node #1"))
	addNode(ctx, fwk, sched, NewTestNode("My node #2"))
	addNode(ctx, fwk, sched, NewTestNode("My node #3"))

	// schedule some pods
	var pods []*v1.Pod
	for i := 0; i < 16; i++ {
		pods = append(pods, NewTestPod(fmt.Sprintf("pod%d", i)))
	}
	for _, pod := range pods {
		scheduler.SchedulePodWithTraits(sched, fwk, &ps, pod, podtraits.AffectNodeCount{})
	}

	return []pretender.StateSnapshot{ps.GetSnapshot()}
}

func main() {
	InitLogs()

	fmt.Println(EvalSchedulerDemo()[0]["my node"])
}
