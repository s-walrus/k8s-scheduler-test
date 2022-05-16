package requests

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/sit/execution"
)

type AddNode struct {
	node *v1.Node
	time int64
}

func (r AddNode) Accept(handler *execution.RequestHandler) error {
	err := handler.UpdateTime(r.time)
	if err != nil {
		return err
	}
	return handler.AddNode(r.node)
}

func NewAddNode(node *v1.Node, time int64) *AddNode {
	return &AddNode{node: node, time: time}
}
