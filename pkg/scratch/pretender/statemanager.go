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

func (s *StateManager) GetSnapshot() StateSnapshot {
	return s.ps.GetSnapshot()
}

func (s *StateManager) GetNode(name string) (*v1.Node, error) {
	return s.ps.GetNode(name)
}

func (s *StateManager) GetPod(uid types.UID) (*PodWithTraits, error) {
	pod, ok := s.podCache[uid]
	if !ok {
		return nil, errors.New("no pod with given name")
	}
	return &pod, nil
}

func (s *StateManager) Bind(nodeName string, podUID types.UID) error {
	pwt, ok := s.podCache[podUID]
	if !ok {
		return errors.New("no pod with given UID in cache")
	}
	err := s.ps.Bind(nodeName, podUID, pwt.Traits)
	if err == nil {
		pwt.Pod.Spec.NodeName = nodeName
	}
	return err
}

func (s *StateManager) AddNode(node *v1.Node) error {
	err := s.ps.AddNode(NewNodeState(node))
	if err == nil {
		s.sched.SchedulerCache.AddNode(node)
	}
	return err
}

func (s *StateManager) RemoveNode(nodeName string) error {
	node, err := s.ps.GetNode(nodeName)
	if err != nil {
		return err
	}
	err = s.ps.RemoveNode(nodeName)
	if err != nil {
		return err
	}

	err = s.sched.SchedulerCache.RemoveNode(node)
	// FIXME node info is not deleted from scheduler cache if pods had been on the node
	return err
}

func (s *StateManager) AddOrUpdatePod(pt PodWithTraits) {
	s.podCache[pt.Pod.UID] = pt
}

func (s *StateManager) RemovePod(uid types.UID) error {
	pod, err := s.GetPod(uid)
	if err != nil {
		return err
	}
	err = s.ps.FindAndRemovePod(uid)
	if err != nil {
		return err
	}

	err = s.sched.SchedulerCache.RemovePod(pod.Pod)
	return err
}

func (s *StateManager) SetFramework(framework framework.Framework) error {
	if s.fwk != nil {
		return errors.New("framework cannot be redefined")
	}
	s.fwk = framework
	return nil
}

func (s *StateManager) UpdateTime(time int64) error {
	return s.UpdateTime(time)
}

func NewStateManager(scheduler *scheduler.Scheduler) StateManager {
	return StateManager{
		ps:       NewState(),
		sched:    scheduler,
		fwk:      nil,
		podCache: map[types.UID]PodWithTraits{},
	}
}
