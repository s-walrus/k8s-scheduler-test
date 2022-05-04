package execution

import (
	"context"
	"k8s.io/kubernetes/pkg/scheduler"
	st "k8s.io/kubernetes/pkg/scheduler/testing"
	"k8s.io/kubernetes/pkg/scratch/pretender"
)

func RunSchedulerIsolationTest(plugins []PluginInfo, scenario RequestGenerator) []pretender.StateSnapshot {
	var rpfs []st.RegisterPluginFunc
	for _, plugin := range plugins {
		rpfs = append(rpfs, plugin.RegisterPluginFunc())
	}
	ctx := context.Background()
	snapshot := scheduler.NewSnapshot()
	sched := scheduler.CreateTestScheduler(ctx, snapshot)
	ps := pretender.NewStateManager(sched)
	fwk := scheduler.NewTestFramework(pretender.NewClientset(&ps), snapshot, rpfs)
	err := ps.SetFramework(fwk)
	if err != nil {
		panic(err)
	}

	handler := NewRequestHandler(&ps, fwk, sched)
	var ret []pretender.StateSnapshot

	for req := scenario.NextRequest(); req != nil; req = scenario.NextRequest() {
		err := req.Accept(&handler)
		if err != nil {
			panic(err)
		}
		ret = append(ret, ps.GetSnapshot())
	}

	return ret
}
