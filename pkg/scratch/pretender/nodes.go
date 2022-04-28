package pretender

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	v12 "k8s.io/client-go/applyconfigurations/core/v1"
	clientset "k8s.io/client-go/kubernetes"
)

type Nodes struct {
	ps     *StateManager
	client clientset.Interface
}

func (c Nodes) Create(ctx context.Context, node *v1.Node, opts metav1.CreateOptions) (*v1.Node, error) {
	err := c.ps.AddNode(node)
	//c.sched.SchedulerCache.AddNode(node)
	return node, err
}

func (Nodes) Update(ctx context.Context, node *v1.Node, opts metav1.UpdateOptions) (*v1.Node, error) {
	//TODO implement me
	panic("implement me")
}

func (Nodes) UpdateStatus(ctx context.Context, node *v1.Node, opts metav1.UpdateOptions) (*v1.Node, error) {
	//TODO implement me
	panic("implement me")
}

func (Nodes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	//TODO implement me
	panic("implement me")
}

func (Nodes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	//TODO implement me
	panic("implement me")
}

func (c Nodes) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Node, error) {
	node, err := c.ps.GetNode(name)
	return node, err
}

func (Nodes) List(ctx context.Context, opts metav1.ListOptions) (*v1.NodeList, error) {
	//TODO implement me
	panic("implement me")
}

func (Nodes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	//TODO implement me
	panic("implement me")
}

func (Nodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Node, err error) {
	//TODO implement me
	panic("implement me")
}

func (Nodes) Apply(ctx context.Context, node *v12.NodeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Node, err error) {
	//TODO implement me
	panic("implement me")
}

func (Nodes) ApplyStatus(ctx context.Context, node *v12.NodeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Node, err error) {
	//TODO implement me
	panic("implement me")
}

func (Nodes) PatchStatus(ctx context.Context, nodeName string, data []byte) (*v1.Node, error) {
	//TODO implement me
	panic("implement me")
}

func NewNodes(ps *StateManager) *Nodes {
	return &Nodes{
		ps: ps,
	}
}
