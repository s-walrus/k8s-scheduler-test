package pretender

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/scheduler"
)

// StateManager is a socket to State, responsible for keeping different subsystems aware of its changes
type StateManager struct {
	ps    State
	sched *scheduler.Scheduler
}

func (c *StateManager) GetSnapshot() StateSnapshot {
	return c.ps.GetSnapshot()
}

func (c *StateManager) GetNode(name string) (*v1.Node, error) {
	return c.ps.GetNode(name)
}

func (c *StateManager) PrepareTraits(traits []PodTrait) bool {
	return c.ps.PrepareTraits(traits)
}

func (c *StateManager) PopPreparedTraits() []PodTrait {
	return c.ps.PopPreparedTraits()
}

func (c *StateManager) Bind(nodeName string, podUID types.UID) error {
	err := c.ps.Bind(nodeName, podUID)
	//if err == nil {
	//	c.sched.SchedulerCache.
	//}
	return err
}

func (c *StateManager) AddNode(node *v1.Node) error {
	err := c.ps.AddNode(NewNodeState(node))
	if err == nil {
		c.sched.SchedulerCache.AddNode(node)
	}
	return err
}

func NewStateManager(scheduler *scheduler.Scheduler) StateManager {
	return StateManager{
		ps:    NewState(),
		sched: scheduler,
	}
}
