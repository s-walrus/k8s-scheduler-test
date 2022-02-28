package pretender

import (
	"k8s.io/client-go/discovery"
	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	admissionregistrationv1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	internalv1alpha1 "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	appsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	authenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	authenticationv1beta1 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	authorizationv1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	autoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	autoscalingv2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2"
	autoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	autoscalingv2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	batchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	certificatesv1 "k8s.io/client-go/kubernetes/typed/certificates/v1"
	"k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	coordinationv1 "k8s.io/client-go/kubernetes/typed/coordination/v1"
	coordinationv1beta1 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	discoveryv1 "k8s.io/client-go/kubernetes/typed/discovery/v1"
	discoveryv1beta1 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
	eventsv1 "k8s.io/client-go/kubernetes/typed/events/v1"
	eventsv1beta1 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	flowcontrolv1alpha1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1"
	flowcontrolv1beta1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
	flowcontrolclient "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2"
	networkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	networkingv1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	nodev1 "k8s.io/client-go/kubernetes/typed/node/v1"
	nodev1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	nodev1beta1 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	policyv1 "k8s.io/client-go/kubernetes/typed/policy/v1"
	policyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	rbacv1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	rbacv1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	schedulingv1 "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	schedulingv1alpha1 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	schedulingv1beta1 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	storagev1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	storagev1alpha1 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	storagev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
)

type Clientset struct {
	core *CoreV1
}

func (c Clientset) Discovery() discovery.DiscoveryInterface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) AdmissionregistrationV1() admissionregistrationv1.AdmissionregistrationV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) AdmissionregistrationV1beta1() admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) InternalV1alpha1() internalv1alpha1.InternalV1alpha1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) AppsV1() appsv1.AppsV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) AppsV1beta1() appsv1beta1.AppsV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) AppsV1beta2() appsv1beta2.AppsV1beta2Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) AuthenticationV1() authenticationv1.AuthenticationV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) AuthenticationV1beta1() authenticationv1beta1.AuthenticationV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) AuthorizationV1() authorizationv1.AuthorizationV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) AuthorizationV1beta1() authorizationv1beta1.AuthorizationV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) AutoscalingV1() autoscalingv1.AutoscalingV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) AutoscalingV2() autoscalingv2.AutoscalingV2Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) AutoscalingV2beta2() autoscalingv2beta2.AutoscalingV2beta2Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) BatchV1() batchv1.BatchV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) BatchV1beta1() batchv1beta1.BatchV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) CertificatesV1() certificatesv1.CertificatesV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) CertificatesV1beta1() v1beta1.CertificatesV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) CoordinationV1beta1() coordinationv1beta1.CoordinationV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) CoordinationV1() coordinationv1.CoordinationV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) CoreV1() corev1.CoreV1Interface {
	return c.core
}

func (c Clientset) DiscoveryV1() discoveryv1.DiscoveryV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) DiscoveryV1beta1() discoveryv1beta1.DiscoveryV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) EventsV1() eventsv1.EventsV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) EventsV1beta1() eventsv1beta1.EventsV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) FlowcontrolV1alpha1() flowcontrolv1alpha1.FlowcontrolV1alpha1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) FlowcontrolV1beta1() flowcontrolv1beta1.FlowcontrolV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) FlowcontrolV1beta2() flowcontrolclient.FlowcontrolV1beta2Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) NetworkingV1() networkingv1.NetworkingV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) NodeV1() nodev1.NodeV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) NodeV1alpha1() nodev1alpha1.NodeV1alpha1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) NodeV1beta1() nodev1beta1.NodeV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) PolicyV1() policyv1.PolicyV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) RbacV1() rbacv1.RbacV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) RbacV1beta1() rbacv1beta1.RbacV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) SchedulingV1() schedulingv1.SchedulingV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) StorageV1beta1() storagev1beta1.StorageV1beta1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) StorageV1() storagev1.StorageV1Interface {
	//TODO implement me
	panic("implement me")
}

func (c Clientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	//TODO implement me
	panic("implement me")
}

func NewPretenderClientset() *Clientset {
	return &Clientset{NewPretenderCoreV1()}
}
