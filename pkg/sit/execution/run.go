package execution

import (
	"context"
	"k8s.io/kubernetes/pkg/scheduler"
	st "k8s.io/kubernetes/pkg/scheduler/testing"
	"k8s.io/kubernetes/pkg/sit/core"
)

func RunSchedulerIsolationTest(plugins []PluginInfo, scenario RequestGenerator) []core.StateSnapshot {
	var rpfs []st.RegisterPluginFunc
	for _, plugin := range plugins {
		rpfs = append(rpfs, plugin.RegisterPluginFunc())
	}
	ctx := context.Background()
	snapshot := scheduler.NewSnapshot()
	sched := scheduler.CreateTestScheduler(ctx, snapshot)
	ps := core.NewStateManager(sched)
	fwk := scheduler.NewTestFramework(core.NewClientset(&ps), snapshot, rpfs)
	err := ps.SetFramework(fwk)
	if err != nil {
		panic(err)
	}

	handler := NewRequestHandler(&ps, fwk, sched)
	var ret []core.StateSnapshot

	for req := scenario.NextRequest(); req != nil; req = scenario.NextRequest() {
		err := req.Accept(&handler)
		if err != nil {
			panic(err)
		}
		ret = append(ret, ps.GetSnapshot())
	}

	return ret
}
