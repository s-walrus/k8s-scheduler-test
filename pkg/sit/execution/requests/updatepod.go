package requests

import (
	"k8s.io/kubernetes/pkg/sit/core"
	"k8s.io/kubernetes/pkg/sit/execution"
)

type UpdatePod struct {
	pod  core.PodWithTraits
	time int64
}

func (r UpdatePod) Accept(handler *execution.RequestHandler) error {
	err := handler.UpdateTime(r.time)
	if err != nil {
		return err
	}
	return handler.UpdatePod(r.pod)
}

func NewUpdatePod(pod core.PodWithTraits, time int64) *UpdatePod {
	return &UpdatePod{pod: pod, time: time}
}
