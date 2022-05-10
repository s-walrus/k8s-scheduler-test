package pretender

import (
	"errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/types"
)

type PodTrait interface {
	Apply(snapshot *NodeSnapshot)
}

type PodWithTraits struct {
	Pod    *v1.Pod
	Traits []PodTrait
}

type nodeState struct {
	v1Node      *v1.Node
	name        string
	cpuCapacity resource.Quantity
	memCapacity resource.Quantity
	pods        map[types.UID][]PodTrait
}

func newNodeState(v1Node *v1.Node, name string, cpu, mem *resource.Quantity) *nodeState {
	return &nodeState{
		v1Node:      v1Node,
		name:        name,
		cpuCapacity: *cpu,
		memCapacity: *mem,
		pods:        make(map[types.UID][]PodTrait),
	}
}

func NewNodeState(node *v1.Node) *nodeState {
	return newNodeState(
		node,
		node.Name,
		node.Status.Capacity.Cpu(),
		node.Status.Capacity.Memory(),
	)
}

type State struct {
	nodes map[string]*nodeState
}

func (s *State) GetNodeSnapshot(nodeName string) (*NodeSnapshot, error) {
	ns, err := s.GetNodeState(nodeName)
	if err != nil {
		return nil, err
	}
	return makeNodeSnapshot(ns), nil
}

func (s *State) GetSnapshot() StateSnapshot {
	snapshot := StateSnapshot{}
	for nodeName, nodeState := range s.nodes {
		snapshot[nodeName] = makeNodeSnapshot(nodeState)
	}
	return snapshot
}

func (s *State) GetNodeState(name string) (*nodeState, error) {
	node, prs := s.nodes[name]
	if !prs {
		return nil, errors.New("no node with name '" + name + "' found")
	}
	return node, nil
}

func (s *State) GetNode(name string) (*v1.Node, error) {
	ns, err := s.GetNodeState(name)
	if err != nil {
		return nil, err
	}
	return ns.v1Node, nil
}

func (s *State) Bind(nodeName string, podUID types.UID, traits []PodTrait) error {
	node, ok := s.nodes[nodeName]
	if !ok {
		return errors.New("no node with name '" + nodeName + "' found")
	}

	node.pods[podUID] = traits
	return nil
}

func (s *State) AddNode(node *nodeState) error {
	// FIXME unnecessary double map lookup
	_, ok := s.nodes[node.name]
	if ok {
		return errors.New("node with given name has been already defined")
	}

	s.nodes[node.name] = node
	return nil
}

func (s *State) RemoveNode(nodeName string) error {
	_, ok := s.nodes[nodeName]
	if !ok {
		return errors.New("no node with given name")
	}

	delete(s.nodes, nodeName)
	return nil
}

func (s *State) FindAndRemovePod(uid types.UID) error {
	for _, node := range s.nodes {
		_, ok := node.pods[uid]
		if ok {
			delete(node.pods, uid)
			return nil
		}
	}
	return errors.New("pod was not found")
}

func NewState() State {
	return State{
		nodes: make(map[string]*nodeState),
	}
}
