package pretendercore

import (
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

type CoreV1 struct{}

func (CoreV1) RESTClient() rest.Interface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) ComponentStatuses() v1.ComponentStatusInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) ConfigMaps(namespace string) v1.ConfigMapInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) Endpoints(namespace string) v1.EndpointsInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) Events(namespace string) v1.EventInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) LimitRanges(namespace string) v1.LimitRangeInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) Namespaces() v1.NamespaceInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) Nodes() v1.NodeInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) PersistentVolumes() v1.PersistentVolumeInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) PersistentVolumeClaims(namespace string) v1.PersistentVolumeClaimInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) Pods(namespace string) v1.PodInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) PodTemplates(namespace string) v1.PodTemplateInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) ReplicationControllers(namespace string) v1.ReplicationControllerInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) ResourceQuotas(namespace string) v1.ResourceQuotaInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) Secrets(namespace string) v1.SecretInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) Services(namespace string) v1.ServiceInterface {
	//TODO implement me
	panic("implement me")
}

func (CoreV1) ServiceAccounts(namespace string) v1.ServiceAccountInterface {
	//TODO implement me
	panic("implement me")
}
