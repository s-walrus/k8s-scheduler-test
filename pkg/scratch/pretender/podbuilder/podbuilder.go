package podbuilder

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/kubernetes/pkg/scratch/pretender"
	"k8s.io/kubernetes/pkg/scratch/pretender/podtraits"
)

type PodBuilder struct {
	pod          pretender.PodWithTraits
	builtPodsCnt int
}

func (c *PodBuilder) GetPod() pretender.PodWithTraits {
	podClone := c.pod.Pod.DeepCopy()
	podName := fmt.Sprintf("%s%d", podClone.ObjectMeta.Name, c.builtPodsCnt)
	podClone.ObjectMeta.Name = podName
	podClone.ObjectMeta.UID = uuid.NewUUID()
	c.builtPodsCnt++
	return pretender.PodWithTraits{
		Pod:    podClone,
		Traits: c.pod.Traits,
	}
}

func (c *PodBuilder) SetMemoryRequest(value int64) *PodBuilder {
	c.pod.Pod.Spec.Containers[0].Resources.Requests[v1.ResourceMemory] = *resource.NewQuantity(value, resource.DecimalSI)
	c.pod.Traits = append(c.pod.Traits, podtraits.RequestMemory{Request: value})
	return c
}

func (c *PodBuilder) SetCPURequest(value int64) *PodBuilder {
	c.pod.Pod.Spec.Containers[0].Resources.Requests[v1.ResourceCPU] = *resource.NewQuantity(value, resource.DecimalSI)
	c.pod.Traits = append(c.pod.Traits, podtraits.RequestCPU{Request: value})
	return c
}

func (c *PodBuilder) AddRequiredPodAntiAffinity(matchLabels map[string]string) *PodBuilder {
	c.pod.Pod.Spec.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution =
		append(c.pod.Pod.Spec.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution,
			newPodAffinityTerm(matchLabels))
	return c
}

func (c *PodBuilder) AddPreferredPodAntiAffinity(matchLabels map[string]string) *PodBuilder {
	c.pod.Pod.Spec.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution =
		append(c.pod.Pod.Spec.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution,
			v1.WeightedPodAffinityTerm{
				Weight:          100,
				PodAffinityTerm: newPodAffinityTerm(matchLabels),
			},
		)
	return c
}

func (c *PodBuilder) AddRequiredPodAffinity(matchLabels map[string]string) *PodBuilder {
	c.pod.Pod.Spec.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution =
		append(c.pod.Pod.Spec.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution,
			newPodAffinityTerm(matchLabels))
	return c
}

func (c *PodBuilder) AddPreferredPodAffinity(matchLabels map[string]string) *PodBuilder {
	c.pod.Pod.Spec.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution =
		append(c.pod.Pod.Spec.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution,
			v1.WeightedPodAffinityTerm{
				Weight:          100,
				PodAffinityTerm: newPodAffinityTerm(matchLabels),
			},
		)
	return c
}

func (c *PodBuilder) SetLabel(name, value string) *PodBuilder {
	c.pod.Pod.ObjectMeta.Labels[name] = value
	return c
}

func NewPodBuilder(name string) *PodBuilder {
	return &PodBuilder{
		pod: pretender.PodWithTraits{
			Pod: &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					UID:       "",
					Namespace: "global-namespace",
					Labels: map[string]string{
						"name": name,
					},
				},
				Spec: v1.PodSpec{
					Affinity: &v1.Affinity{
						NodeAffinity:    &v1.NodeAffinity{},
						PodAffinity:     &v1.PodAffinity{},
						PodAntiAffinity: &v1.PodAntiAffinity{},
					},
					Containers: []v1.Container{
						{
							Resources: v1.ResourceRequirements{
								Limits:   nil,
								Requests: v1.ResourceList{},
							},
						},
					},
				},
			},
			Traits: []pretender.PodTrait{
				podtraits.AffectNodeCount{},
			},
		},
	}
}

func newPodAffinityTerm(matchLabels map[string]string) v1.PodAffinityTerm {
	return v1.PodAffinityTerm{
		LabelSelector: &metav1.LabelSelector{
			MatchLabels:      matchLabels,
			MatchExpressions: []metav1.LabelSelectorRequirement{},
		},
		Namespaces:        []string{"global-namespace"},
		TopologyKey:       "name",
		NamespaceSelector: nil,
	}
}
