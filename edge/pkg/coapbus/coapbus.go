package coapbus

import (
	"k8s.io/klog"

	"github.com/kubeedge/beehive/pkg/core"
	beehiveContext "github.com/kubeedge/beehive/pkg/core/context"
	coapconfig "github.com/kubeedge/kubeedge/edge/pkg/coapbus/config"
	"github.com/kubeedge/kubeedge/edge/pkg/common/modules"
	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/edgecore/v1alpha1"
)

// eventbus struct
type coapbus struct {
	enable bool
}

func newCoapbus(enable bool) *coapbus {
	return &coapbus{
		enable: enable,
	}
}

// Register register coapbus
func Register(coapbus *v1alpha1.CoapBus, nodeName string) {
	coapconfig.InitConfigure(coapbus, nodeName)
	core.Register(newCoapbus(coapbus.Enable))
}

func (*coapbus) Name() string {
	return "coapbus"
}

func (*coapbus) Group() string {
	return modules.CoapGroup
}

// Enable indicates wheather this module is enabled
func (cb *coapbus) Enable() bool {
	return cb.enable
}

func (cb *coapbus) Start() {
	klog.Info("CoapBus inside the start function")
	cb.pubCloudMsgToEdge()
}

func (cb *coapbus) pubCloudMsgToEdge() {
	for {
		select {
		case <-beehiveContext.Done():
			klog.Warning("CoapBus PubCloudMsg to Edge stop")
			return
		default:
		}
	}
}
