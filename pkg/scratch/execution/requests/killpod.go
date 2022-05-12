package requests

import (
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/scratch/execution"
)

type KillPod struct {
	uid  types.UID
	time int64
}

func (r KillPod) Accept(handler *execution.RequestHandler) error {
	err := handler.UpdateTime(r.time)
	if err != nil {
		return err
	}
	return handler.KillPod(r.uid)
}

func NewKillPod(uid types.UID, time int64) *KillPod {
	return &KillPod{uid: uid, time: time}
}
