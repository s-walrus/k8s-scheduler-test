package pretender

import (
	"errors"
	"k8s.io/apimachinery/pkg/types"
)

type PodsSnapshot struct {
}

type nodeState struct {
	name        string
	cpuCapacity float64
	memCapacity float64
	pods        map[types.UID][]*PodTrait
}

func NewNodeState(name string, cpu, mem float64) *nodeState {
	return &nodeState{
		name:        name,
		cpuCapacity: cpu,
		memCapacity: mem,
		pods:        make(map[types.UID][]*PodTrait),
	}
}

type State struct {
	nodes          map[string]*nodeState
	preparedTraits []*PodTrait
}

func (c *State) GetSnapshot() PodsSnapshot {
	return PodsSnapshot{
		// TODO
	}
}

func (c *State) PrepareTraits(traits []*PodTrait) bool {
	ok := c.preparedTraits == nil
	c.preparedTraits = traits
	return ok
}

func (c *State) PopPreparedTraits() []*PodTrait {
	ret := c.preparedTraits
	c.preparedTraits = nil
	return ret
}

func (c *State) Bind(nodeName string, podUID types.UID) error {
	traits := c.PopPreparedTraits()
	if traits == nil {
		return errors.New("traits must be prepared before binding")
	}

	node, prs := c.nodes[nodeName]
	if !prs {
		return errors.New("node with given name was not initialized")
	}

	// assuming all UIDs are unique
	node.pods[podUID] = traits

	return nil
}

func (c *State) AddNode(node *nodeState) error {
	// FIXME unnecessary double map lookup
	_, prs := c.nodes[node.name]
	if prs {
		return errors.New("node with given name has been already defined")
	}

	c.nodes[node.name] = node
	return nil
}

func NewState() State {
	return State{
		nodes:          make(map[string]*nodeState),
		preparedTraits: nil,
	}
}
