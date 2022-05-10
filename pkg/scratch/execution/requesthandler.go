package execution

import (
	"context"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/scheduler"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scratch/pretender"
)

// RequestHandler runs scheduling logic by request from Request class
// (may require refactoring to better reflect its purpose)
type RequestHandler struct {
	ps    *pretender.StateManager
	fwk   framework.Framework
	sched *scheduler.Scheduler
}

func (c *RequestHandler) AddNode(node *v1.Node) error {
	_, err := c.fwk.ClientSet().CoreV1().Nodes().Create(context.Background(), node, metav1.CreateOptions{})
	return err
}

func (c *RequestHandler) RemoveNode(name string) error {
	err := c.fwk.ClientSet().CoreV1().Nodes().Delete(context.Background(), name, metav1.DeleteOptions{})
	return err
}

func (c *RequestHandler) SchedulePod(pod pretender.PodWithTraits) error {
	c.ps.AddOrUpdatePod(pod)
	scheduler.FakeScheduleOne(context.Background(), c.sched, c.fwk, pod.Pod)
	return nil
}

func (c *RequestHandler) KillPod(uid types.UID) error {
	err := c.fwk.ClientSet().CoreV1().Pods("").EvictV1(context.Background(), &v12.Eviction{ObjectMeta: metav1.ObjectMeta{UID: uid}})
	return err
}

func NewRequestHandler(ps *pretender.StateManager, fwk framework.Framework, sched *scheduler.Scheduler) RequestHandler {
	return RequestHandler{
		ps:    ps,
		fwk:   fwk,
		sched: sched,
	}
}
