package requests

import (
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/scratch/execution"
)

type KillPod struct {
	uid types.UID
}

func (r KillPod) Accept(handler *execution.RequestHandler) error {
	return handler.KillPod(r.uid)
}

func NewKillPod(uid types.UID) *KillPod {
	return &KillPod{uid: uid}
}
