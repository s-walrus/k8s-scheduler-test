package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler"
	"k8s.io/kubernetes/pkg/scratch/pretender"
)

func InitLogs() {
	klog.InitFlags(nil)
	flag.Parse()
}

func EvalSchedulerDemo() []pretender.PodsSnapshot {
	ctx := context.Background()
	ps := pretender.NewState()
	sched := scheduler.CreateTestScheduler(ctx)
	fwk := scheduler.NewTestFramework(&ps)

	err := ps.AddNode(pretender.NewNodeState("my node", 0, 0))
	if err != nil {
		panic(err)
	}

	// schedule some pods
	pod := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "123", UID: types.UID("321")}, Spec: v1.PodSpec{}}
	scheduler.SchedulePodWithTraits(sched, fwk, &ps, pod, []*pretender.PodTrait{})

	return []pretender.PodsSnapshot{ps.GetSnapshot()}
}

func main() {
	InitLogs()

	fmt.Println(EvalSchedulerDemo()[0])
}
