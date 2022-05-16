package requests

import (
	"k8s.io/kubernetes/pkg/sit/execution"
)

type RemoveNode struct {
	name string
	time int64
}

func (r RemoveNode) Accept(handler *execution.RequestHandler) error {
	err := handler.UpdateTime(r.time)
	if err != nil {
		return err
	}
	return handler.RemoveNode(r.name)
}

func NewRemoveNode(name string, time int64) *RemoveNode {
	return &RemoveNode{name: name, time: time}
}
