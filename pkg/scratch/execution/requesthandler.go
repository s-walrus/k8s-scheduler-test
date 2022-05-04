package execution

import (
	"k8s.io/kubernetes/pkg/scheduler"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scratch/pretender"
)

// RequestHandler runs scheduling logic by request from Request class
// (may require refactoring to better reflect its purpose)
type RequestHandler struct {
	ps    *pretender.StateManager
	fwk   framework.Framework
	sched *scheduler.Scheduler
}

// here go different requests handlers...

func NewRequestHandler(ps *pretender.StateManager, fwk framework.Framework, sched *scheduler.Scheduler) RequestHandler {
	return RequestHandler{
		ps:    ps,
		fwk:   fwk,
		sched: sched,
	}
}
