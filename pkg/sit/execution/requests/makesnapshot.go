package requests

import (
	"k8s.io/kubernetes/pkg/sit/execution"
)

type MakeSnapshot struct {
	time int64
}

func (r MakeSnapshot) Accept(handler *execution.RequestHandler) error {
	return handler.UpdateTime(r.time)
}

func NewMakeSnapshot(time int64) *MakeSnapshot {
	return &MakeSnapshot{time: time}
}
