package requests

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scratch/execution"
)

type AddNode struct {
	node *v1.Node
}

func (r AddNode) Accept(handler *execution.RequestHandler) error {
	return handler.AddNode(r.node)
}

func NewAddNode(node *v1.Node) *AddNode {
	return &AddNode{node: node}
}
