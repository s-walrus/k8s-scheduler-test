package main

import (
	"context"
	"flag"
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

	// schedule some pods
	pod := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "123", UID: types.UID("321")}, Spec: v1.PodSpec{}}
	scheduler.FakeScheduleOne(ctx, sched, fwk, pod)

	return []pretender.PodsSnapshot{ps.GetSnapshot()}
}

func main() {
	InitLogs()

	EvalSchedulerDemo()
}
