package scenarios

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/kubernetes/pkg/scratch/execution"
	"k8s.io/kubernetes/pkg/scratch/execution/requests"
	"k8s.io/kubernetes/pkg/scratch/pretender/podbuilder"
)

// TODO add node builder, remove NewTestNode

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
				v1.ResourceCPU:    *(resource.NewQuantity(100500, resource.DecimalSI)),
				v1.ResourceMemory: *(resource.NewQuantity(100500, resource.DecimalSI)),
				v1.ResourcePods:   *(resource.NewQuantity(10, resource.DecimalSI)),
			},
			Allocatable: v1.ResourceList{
				v1.ResourceCPU:    *(resource.NewQuantity(100500, resource.DecimalSI)),
				v1.ResourceMemory: *(resource.NewQuantity(100500, resource.DecimalSI)),
				v1.ResourcePods:   *(resource.NewQuantity(10, resource.DecimalSI)),
			}},
	}
	return &node
}

func NewSelfAntiAffinityPods() *execution.StaticRequestGenerator {
	var reqs []execution.Request
	for i := 0; i < 3; i++ {
		reqs = append(reqs, requests.NewAddNode(NewTestNode(fmt.Sprintf("node%d", i+1))))
	}

	affinityPodBuilder := podbuilder.NewPodBuilder("affinity-pod")
	affinityPodBuilder.AddPreferredPodAntiAffinity(map[string]string{"affinity-group": "1"})
	affinityPodBuilder.SetLabel("affinity-group", "1")
	for i := 0; i < 16; i++ {
		reqs = append(reqs, requests.NewSchedulePod(affinityPodBuilder.GetPod()))
	}
	return execution.NewStaticRequestGenerator(reqs)
}
