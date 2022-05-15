package main

import (
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/defaultbinder"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/feature"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/interpodaffinity"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/noderesources"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/queuesort"
	st "k8s.io/kubernetes/pkg/scheduler/testing"
	"k8s.io/kubernetes/pkg/scratch/execution"
	"k8s.io/kubernetes/pkg/scratch/execution/requests"
	"k8s.io/kubernetes/pkg/scratch/pretender"
	"k8s.io/kubernetes/pkg/scratch/pretender/podbuilder"
	"k8s.io/kubernetes/pkg/scratch/pretender/podtraits"
	"math"
	"math/rand"
	"time"
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
				"\t%s: { cnt: %d, mem: %d, cpu: %d, rcpu: %f, t: %d }\n",
				node,
				state.NodeCount,
				state.MemoryRequested,
				state.MilliCPURequested,
				state.CPULoad,
				state.Time,
			)
		}
		fmt.Println("},")
	}
}

func CountCPUOverloadTime(snapshots []pretender.StateSnapshot, nodeCapacity map[string]int64) map[string]int64 {
	olTime := map[string]int64{}

	for i, snapshot := range snapshots {
		if i == 0 {
			continue
		}
		for k, ns := range snapshot {
			if int64(ns.CPULoad) >= nodeCapacity[k] {
				prevTime := snapshots[i-1][k].Time
				olTime[k] += ns.Time - prevTime
			}
		}
	}

	return olTime
}

func GetAverageCPULoadTimeRatio(snapshots []pretender.StateSnapshot, nodeCapacity map[string]int64) float64 {
	var nodes []string
	for k := range nodeCapacity {
		nodes = append(nodes, k)
	}

	overloadTime := CountCPUOverloadTime(snapshots, nodeCapacity)
	timeTotal := snapshots[len(snapshots)-1][nodes[0]].Time - 0

	var sum int64 = 0
	for _, v := range overloadTime {
		sum += v
	}
	ratio := float64(sum) / float64(timeTotal) / float64(len(nodes))

	return ratio
}

// FIXME add node builder, remove NewTestNode

func NewTestNode(name string) *v1.Node {
	node := v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			UID:  uuid.NewUUID(),
			Labels: map[string]string{
				"name": name,
			},
		},
		Status: v1.NodeStatus{
			Capacity: v1.ResourceList{
				v1.ResourceCPU:    *(resource.NewQuantity(0, resource.DecimalSI)),
				v1.ResourceMemory: *(resource.NewQuantity(0, resource.DecimalSI)),
				v1.ResourcePods:   *(resource.NewQuantity(0, resource.DecimalSI)),
			},
			Allocatable: v1.ResourceList{
				v1.ResourceCPU:    *(resource.NewQuantity(0, resource.DecimalSI)),
				v1.ResourceMemory: *(resource.NewQuantity(0, resource.DecimalSI)),
				v1.ResourcePods:   *(resource.NewQuantity(0, resource.DecimalSI)),
			}},
	}
	return &node
}

func NewRandomPodBuilder(accurateEstimate bool) (*podbuilder.PodBuilder, *podtraits.FiniteFourierSeries) {
	var sinKs, cosKs []float64
	var sumAbs float64 = 0
	for i := 0; i < 8; i++ {
		sinKs = append(sinKs, 1000*(rand.Float64()*2-1)/float64(i+4))
		sumAbs += math.Abs(sinKs[i])
	}
	for i := 0; i < 8; i++ {
		cosKs = append(cosKs, 1000*(rand.Float64()*2-1)/float64(i+4))
		sumAbs += math.Abs(cosKs[i])
	}
	k := sumAbs / 2

	cpuFunc := podtraits.NewFiniteFourierSeries(k, sinKs, cosKs)
	var cpuEstimate float64
	if accurateEstimate {
		cpuEstimate = k
	} else {
		cpuEstimate = cpuFunc.Integrate(-1000, 0)
	}

	pb := podbuilder.NewPodBuilder(fmt.Sprintf("random-pod-%d", rand.Intn(1000)))
	pb.AddCPUUsageFunc(podtraits.NewFiniteFourierSeries(k, sinKs, cosKs))
	pb.SetCPURequest(int64(cpuEstimate))
	return pb, cpuFunc
}

func MyTestScenario(updateRequests bool, accurateEstimate bool) *execution.StaticRequestGenerator {
	var reqs []execution.Request
	for i := 0; i < 4; i++ {
		reqs = append(reqs, requests.NewAddNode(NewTestNode(fmt.Sprintf("node%d", i+1)), 0))
	}

	var pods []pretender.PodWithTraits
	var cpuFuncs []*podtraits.FiniteFourierSeries

	for i := 0; i < 10000; i++ {
		if i%100 == 0 {
			if updateRequests {
				for i, pod := range pods {
					cpuEstimate := cpuFuncs[i].Integrate(float64((i-1)*1000), float64(i*1000))
					pod.Pod.Spec.Containers[0].Resources.Requests[v1.ResourceCPU] = *resource.NewQuantity(int64(cpuEstimate), resource.DecimalSI)
				}
			}
			pb, cpuFunc := NewRandomPodBuilder(accurateEstimate)
			pod := pb.GetPod()
			pods = append(pods, pod)
			cpuFuncs = append(cpuFuncs, cpuFunc)
			reqs = append(reqs, requests.NewSchedulePod(pod, int64(i*1000)))
		}
		reqs = append(reqs, requests.NewMakeSnapshot(int64(i*1000)))
	}
	return execution.NewStaticRequestGenerator(reqs)
}

func main() {
	InitLogs()
	rand.Seed(int64(time.Now().Second()))

	plugins := []execution.PluginInfo{
		execution.NewPluginInfo(queuesort.Name, queuesort.New, "QueueSort"),
		execution.NewPluginInfo(interpodaffinity.Name, NewInterPodAffinity, "PreFilter", "Filter", "PreScore", "Score"),
		execution.NewPluginInfo(noderesources.BalancedAllocationName, NewBalancedAllocation, "Score"),
		execution.NewPluginInfo("TrueFilter", st.NewTrueFilterPlugin, "Filter"),
		execution.NewPluginInfo(defaultbinder.Name, defaultbinder.New, "Bind"),
	}

	nodeCapacity := map[string]int64{
		"node1": 15000,
		"node2": 15000,
		"node3": 15000,
		"node4": 15000,
	}

	// inaccurate cpu load estimate, no cpu load updates
	{
		for i := 0; i < 1000; i++ {
			snapshots := execution.RunSchedulerIsolationTest(plugins, MyTestScenario(false, false))
			fmt.Println(GetAverageCPULoadTimeRatio(snapshots, nodeCapacity))
		}
	}

	// accurate estimate, no update
	{
		for i := 0; i < 1000; i++ {
			snapshots := execution.RunSchedulerIsolationTest(plugins, MyTestScenario(false, true))
			fmt.Println(GetAverageCPULoadTimeRatio(snapshots, nodeCapacity))
		}
	}

	// inaccurate estimate, with updates
	{
		for i := 0; i < 1000; i++ {
			snapshots := execution.RunSchedulerIsolationTest(plugins, MyTestScenario(true, false))
			fmt.Println(GetAverageCPULoadTimeRatio(snapshots, nodeCapacity))
		}
	}
}

/*

Closest project goals:
+ test if updating pod resource requests is beneficial
+ [come up with a plugin and test it on a relatively large scale]

Tasks:
$ add a pod trait that reflects real resource usage
	$ figure out how to do it with basic fourier series
	$ add time consideration to pod traits handling
	$ implement the pod trait
	$ track choke time
$ create a scenario with pods consuming random real resources
	$ implement a pod builder for random resource consuming pods (random are fourier series coefficients [realistic though])
	* add an option to make synchronized spikes in resource usage
	$ add "request update" request
	$ create the scenario and a similar one with request updates
* run the test with different configurations of the scenario (similar function shifted, for example)
* make some useful metrics from test results
? come up with a plugin and run the implemented scenario with it

*/
