package core

import (
	"context"
	v1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	"k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Pods struct {
	ps     *StateManager
	client clientset.Interface // FIXME do i need clientset here? also in Nodes
}

func (p Pods) Create(ctx context.Context, pod *v1.Pod, opts metav1.CreateOptions) (*v1.Pod, error) {
	//TODO implement me
	panic("implement me")
}

func (p Pods) Update(ctx context.Context, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error) {
	//TODO implement me
	panic("implement me")
}

func (p Pods) UpdateStatus(ctx context.Context, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error) {
	//TODO implement me
	panic("implement me")
}

func (p Pods) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	//err := p.ps.RemovePod(name)
	//return err
	//TODO implement me
	panic("implement me")
}

func (p Pods) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	//TODO implement me
	panic("implement me")
}

func (p Pods) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Pod, error) {
	//TODO implement me
	panic("implement me")
}

func (p Pods) List(ctx context.Context, opts metav1.ListOptions) (*v1.PodList, error) {
	//TODO implement me
	panic("implement me")
}

func (p Pods) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	//TODO implement me
	panic("implement me")
}

func (p Pods) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Pod, err error) {
	//TODO implement me
	panic("implement me")
}

func (p Pods) Apply(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Pod, err error) {
	//TODO implement me
	panic("implement me")
}

func (p Pods) ApplyStatus(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Pod, err error) {
	//TODO implement me
	panic("implement me")
}

func (p Pods) UpdateEphemeralContainers(ctx context.Context, podName string, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error) {
	//TODO implement me
	panic("implement me")
}

func (p Pods) Bind(ctx context.Context, binding *v1.Binding, opts metav1.CreateOptions) error {
	//nodeName := binding.Target.Name
	//podName := binding.ObjectMeta.Name
	nodeName := binding.Target.Name
	podUID := binding.ObjectMeta.UID
	err := p.ps.Bind(nodeName, podUID)
	return err
}

func (p Pods) Evict(ctx context.Context, eviction *v1beta1.Eviction) error {
	//TODO implement me
	panic("implement me")
}

func (p Pods) EvictV1(ctx context.Context, eviction *policyv1.Eviction) error {
	return p.ps.RemovePod(eviction.UID)
}

func (p Pods) EvictV1beta1(ctx context.Context, eviction *v1beta1.Eviction) error {
	//TODO implement me
	panic("implement me")
}

func (p Pods) GetLogs(name string, opts *v1.PodLogOptions) *rest.Request {
	//TODO implement me
	panic("implement me")
}

func (p Pods) ProxyGet(scheme, name, port, path string, params map[string]string) rest.ResponseWrapper {
	//TODO implement me
	panic("implement me")
}

func NewPods(ps *StateManager) *Pods {
	return &Pods{
		ps: ps,
	}
}
