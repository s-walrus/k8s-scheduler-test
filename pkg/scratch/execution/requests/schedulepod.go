package requests

import (
	"k8s.io/kubernetes/pkg/scratch/execution"
	"k8s.io/kubernetes/pkg/scratch/pretender"
)

type SchedulePod struct {
	pod pretender.PodWithTraits
}

func (r SchedulePod) Accept(handler *execution.RequestHandler) error {
	return handler.SchedulePod(r.pod)
}

func NewSchedulePod(pod pretender.PodWithTraits) *SchedulePod {
	return &SchedulePod{pod: pod}
}
