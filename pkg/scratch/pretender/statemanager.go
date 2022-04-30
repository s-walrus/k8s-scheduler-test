package pretender

import (
	"errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/scheduler"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// StateManager is a socket to State, responsible for keeping different subsystems aware of its changes
type StateManager struct {
	ps    State
	sched *scheduler.Scheduler
	fwk   framework.Framework

	// FIXME remove unused pods from cache
	podCache map[types.UID]PodWithTraits
}

func (c *StateManager) GetSnapshot() StateSnapshot {
	return c.ps.GetSnapshot()
}

func (c *StateManager) GetNode(name string) (*v1.Node, error) {
	return c.ps.GetNode(name)
}

func (c *StateManager) Bind(nodeName string, podUID types.UID) error {
	pwt, ok := c.podCache[podUID]
	if !ok {
		return errors.New("no pod with given UID in cache")
	}
	err := c.ps.Bind(nodeName, podUID, pwt.Traits)
	if err == nil {
		pwt.Pod.Spec.NodeName = nodeName
		node, err := c.fwk.SnapshotSharedLister().NodeInfos().Get(nodeName)
		if err != nil {
			panic(err)
		}
		node.Requested.MilliCPU = 1
		node.Requested.Memory = 1
	}
	//if err == nil {
	//	node, err := c.fwk.SnapshotSharedLister().NodeInfos().Get(nodeName)
	//	if err != nil {
	//		panic(err)
	//	}
	//	node.Requested.MilliCPU = snapshot.MilliCPURequested
	//	node.Requested.Memory = snapshot.MemoryRequested
	//	err := c.sched.SchedulerCache.AddPod(pod.pod)
	//	if err != nil {
	//		panic(err)
	//	}
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

// AddOrUpdatePod adds pod to cache or updates existing pod with same UID
func (c *StateManager) AddOrUpdatePod(pt PodWithTraits) {
	c.podCache[pt.Pod.UID] = pt
}

func (c *StateManager) SetFramework(framework framework.Framework) error {
	if c.fwk != nil {
		return errors.New("framework cannot be redefined")
	}
	c.fwk = framework
	return nil
}

func NewStateManager(scheduler *scheduler.Scheduler) StateManager {
	return StateManager{
		ps:       NewState(),
		sched:    scheduler,
		fwk:      nil,
		podCache: map[types.UID]PodWithTraits{},
	}
}
