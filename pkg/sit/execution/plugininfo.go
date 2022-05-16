package execution

import (
	"k8s.io/kubernetes/pkg/scheduler/framework/runtime"
	st "k8s.io/kubernetes/pkg/scheduler/testing"
)

type PluginInfo struct {
	pluginName    string
	pluginNewFunc runtime.PluginFactory
	extensions    []string
}

func (c PluginInfo) RegisterPluginFunc() st.RegisterPluginFunc {
	return st.RegisterPluginAsExtensions(c.pluginName, c.pluginNewFunc, c.extensions...)
}

func NewPluginInfo(pluginName string, pluginNewFunc runtime.PluginFactory, extensions ...string) PluginInfo {
	return PluginInfo{
		pluginName:    pluginName,
		pluginNewFunc: pluginNewFunc,
		extensions:    extensions,
	}
}
