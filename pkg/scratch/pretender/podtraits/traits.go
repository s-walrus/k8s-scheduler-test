package podtraits

import (
	"k8s.io/kubernetes/pkg/scratch/pretender"
)

type AffectNodeCount struct{}

func (AffectNodeCount) Apply(snapshot *pretender.NodeSnapshot) {
	snapshot.NodeCount++
}

type RequestMemory struct {
	Request int64
}

func (c RequestMemory) Apply(snapshot *pretender.NodeSnapshot) {
	snapshot.MemoryRequested += c.Request
}

type RequestCPU struct {
	Request int64
}

func (c RequestCPU) Apply(snapshot *pretender.NodeSnapshot) {
	snapshot.MilliCPURequested += c.Request
}
