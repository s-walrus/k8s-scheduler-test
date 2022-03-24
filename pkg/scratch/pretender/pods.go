package pretender

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
	ps     *State
	client clientset.Interface
}

func (c Pods) Create(ctx context.Context, pod *v1.Pod, opts metav1.CreateOptions) (*v1.Pod, error) {
	//TODO implement me
	panic("implement me")
}

func (c Pods) Update(ctx context.Context, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error) {
	//TODO implement me
	panic("implement me")
}

func (c Pods) UpdateStatus(ctx context.Context, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error) {
	//TODO implement me
	panic("implement me")
}

func (c Pods) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	//TODO implement me
	panic("implement me")
}

func (c Pods) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	//TODO implement me
	panic("implement me")
}

func (c Pods) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Pod, error) {
	//TODO implement me
	panic("implement me")
}

func (c Pods) List(ctx context.Context, opts metav1.ListOptions) (*v1.PodList, error) {
	//TODO implement me
	panic("implement me")
}

func (c Pods) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	//TODO implement me
	panic("implement me")
}

func (c Pods) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Pod, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Pods) Apply(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Pod, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Pods) ApplyStatus(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Pod, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Pods) UpdateEphemeralContainers(ctx context.Context, podName string, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error) {
	//TODO implement me
	panic("implement me")
}

func (c Pods) Bind(ctx context.Context, binding *v1.Binding, opts metav1.CreateOptions) error {
	//nodeName := binding.Target.Name
	//podName := binding.ObjectMeta.Name
	nodeName := binding.Target.Name
	podUID := binding.ObjectMeta.UID
	err := c.ps.Bind(nodeName, podUID)
	return err
}

func (c Pods) Evict(ctx context.Context, eviction *v1beta1.Eviction) error {
	//TODO implement me
	panic("implement me")
}

func (c Pods) EvictV1(ctx context.Context, eviction *policyv1.Eviction) error {
	//TODO implement me
	panic("implement me")
}

func (c Pods) EvictV1beta1(ctx context.Context, eviction *v1beta1.Eviction) error {
	//TODO implement me
	panic("implement me")
}

func (c Pods) GetLogs(name string, opts *v1.PodLogOptions) *rest.Request {
	//TODO implement me
	panic("implement me")
}

func (c Pods) ProxyGet(scheme, name, port, path string, params map[string]string) rest.ResponseWrapper {
	//TODO implement me
	panic("implement me")
}

func NewPods(ps *State) *Pods {
	return &Pods{
		ps: ps,
	}
}
