package podtraits

import (
	"k8s.io/kubernetes/pkg/sit/core"
)

type AffectNodeCount struct{}

func (AffectNodeCount) Apply(snapshot *core.NodeSnapshot, _ int64) {
	snapshot.NodeCount++
}

type RequestMemory struct {
	Request int64
}

func (c RequestMemory) Apply(snapshot *core.NodeSnapshot, _ int64) {
	snapshot.MemoryRequested += c.Request
}

type RequestCPU struct {
	Request int64
}

func (t RequestCPU) Apply(snapshot *core.NodeSnapshot, _ int64) {
	snapshot.MilliCPURequested += t.Request
}

type WithComplexCPUUsage struct {
	UsageFunc *FiniteFourierSeries
}

func (t WithComplexCPUUsage) Apply(snapshot *core.NodeSnapshot, time int64) {
	snapshot.CPULoad += t.UsageFunc.GetValue(float64(time) / (1 << 20))
}
