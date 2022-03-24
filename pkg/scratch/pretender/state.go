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
	nodes          map[string]*nodeState
	preparedTraits []PodTrait
}

func (c *State) GetSnapshot() StateSnapshot {
	snapshot := StateSnapshot{}
	for nodeName, nodeState := range c.nodes {
		snapshot[nodeName] = makeNodeSnapshot(nodeState)
	}
	return snapshot
}

func (c *State) GetNode(name string) (*v1.Node, error) {
	node, prs := c.nodes[name]
	if !prs {
		return nil, errors.New("no node with name '" + name + "' found")
	}
	return node.v1Node, nil
}

func (c *State) PrepareTraits(traits []PodTrait) bool {
	ok := c.preparedTraits == nil
	c.preparedTraits = traits
	return ok
}

func (c *State) PopPreparedTraits() []PodTrait {
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
		return errors.New("no node with name '" + nodeName + "' found")
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
