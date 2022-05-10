package main

import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/defaultbinder"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/feature"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/interpodaffinity"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/noderesources"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/queuesort"
	st "k8s.io/kubernetes/pkg/scheduler/testing"
	"k8s.io/kubernetes/pkg/scratch/execution"
	"k8s.io/kubernetes/pkg/scratch/execution/scenarios"
	"k8s.io/kubernetes/pkg/scratch/pretender"
)

func InitLogs() {
	klog.InitFlags(nil)
	flag.Parse()
}

func NewInterPodAffinity(plArgs runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return interpodaffinity.New(plArgs, h, feature.Features{
		EnablePodAffinityNamespaceSelector: true,
		EnablePodDisruptionBudget:          false,
		EnablePodOverhead:                  false,
		EnableReadWriteOncePod:             false,
		EnableVolumeCapacityPriority:       false,
		EnableCSIStorageCapacity:           false,
	})
}

func NewBalancedAllocation(plArgs runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return noderesources.NewBalancedAllocation(plArgs, h, feature.Features{
		EnablePodAffinityNamespaceSelector: true,
		EnablePodDisruptionBudget:          false,
		EnablePodOverhead:                  false,
		EnableReadWriteOncePod:             false,
		EnableVolumeCapacityPriority:       false,
		EnableCSIStorageCapacity:           false,
	})
}

func PrintTestResult(snapshots []pretender.StateSnapshot) {
	for _, s := range snapshots {
		fmt.Println("{")
		for node, state := range s {
			fmt.Printf(
				"\t%s: { cnt: %d, mem: %d, cpu: %d }\n",
				node,
				state.NodeCount,
				state.MemoryRequested,
				state.MilliCPURequested,
			)
		}
		fmt.Println("},")
	}
}

func main() {
	InitLogs()

	plugins := []execution.PluginInfo{
		execution.NewPluginInfo(queuesort.Name, queuesort.New, "QueueSort"),
		execution.NewPluginInfo(interpodaffinity.Name, NewInterPodAffinity, "PreFilter", "Filter", "PreScore", "Score"),
		execution.NewPluginInfo(noderesources.BalancedAllocationName, NewBalancedAllocation, "Score"),
		execution.NewPluginInfo("TrueFilter", st.NewTrueFilterPlugin, "Filter"),
		execution.NewPluginInfo(defaultbinder.Name, defaultbinder.New, "Bind"),
	}

	PrintTestResult(execution.RunSchedulerIsolationTest(plugins, scenarios.SelfAntiAffinityPods()))
}
