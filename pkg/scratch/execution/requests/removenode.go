package requests

import (
	"k8s.io/kubernetes/pkg/scratch/execution"
)

type RemoveNode struct {
	name string
}

func (r RemoveNode) Accept(handler *execution.RequestHandler) error {
	return handler.RemoveNode(r.name)
}

func NewRemoveNode(name string) *RemoveNode {
	return &RemoveNode{name: name}
}
