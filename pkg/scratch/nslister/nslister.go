package staticnslister

import (
	"errors"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type StaticNsLister struct {
	namespaces map[string]*v1.Namespace
}

func (s StaticNsLister) List(_ labels.Selector) ([]*v1.Namespace, error) {
	var ret []*v1.Namespace
	for _, v := range s.namespaces {
		ret = append(ret, v)
	}
	return ret, nil
}

func (s StaticNsLister) Get(name string) (*v1.Namespace, error) {
	ns, isPresent := s.namespaces[name]
	if !isPresent {
		return nil, errors.New("namespace with given name does not exist")
	}
	return ns, nil
}

func NewStaticNsLister(nsNames ...string) *StaticNsLister {
	nsLister := StaticNsLister{
		map[string]*v1.Namespace{},
	}
	for _, name := range nsNames {
		nsLister.namespaces[name] = &v1.Namespace{
			ObjectMeta: v12.ObjectMeta{
				Name: name,
			},
		}
	}
	return &nsLister
}
