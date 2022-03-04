package podtraits

import (
	"k8s.io/kubernetes/pkg/scratch/pretender"
)

type AffectNodeCount struct{}

func (AffectNodeCount) Apply(snapshot *pretender.NodeSnapshot) {
	snapshot.NodeCount++
}
