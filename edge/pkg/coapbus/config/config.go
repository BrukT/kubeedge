package config

import (
	"sync"

	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/edgecore/v1alpha1"
)

var Config Configure
var once sync.Once

type Configure struct {
	v1alpha1.CoapBus
	NodeName string
}

func InitConfigure(coapbus *v1alpha1.CoapBus, nodeName string) {
	once.Do(func() {
		Config = Configure{
			CoapBus:  *coapbus,
			NodeName: nodeName,
		}
	})
}
