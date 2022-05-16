package requests

import (
	"k8s.io/kubernetes/pkg/sit/core"
	"k8s.io/kubernetes/pkg/sit/execution"
)

type SchedulePod struct {
	pod  core.PodWithTraits
	time int64
}

func (r SchedulePod) Accept(handler *execution.RequestHandler) error {
	err := handler.UpdateTime(r.time)
	if err != nil {
		return err
	}
	return handler.SchedulePod(r.pod)
}

func NewSchedulePod(pod core.PodWithTraits, time int64) *SchedulePod {
	return &SchedulePod{pod: pod, time: time}
}
