package test

import (
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"k8s.io/klog"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// Begin starts the udp server
func Begin() {
	PORT := ":5683"

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		klog.Errorf("Failed to resolve the port addr: %v", err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		klog.Errorf("Failed to resolve ip: %v", err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		klog.Infoln("CoapBus", addr, " -> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			klog.Infof("CoapBus Exiting UDP server!")
			return
		}

		data1 := []byte(strconv.Itoa(random(1, 1001)))
		data := []byte("Hi I am the server")
		//swapped the data with data1 from the original just to chek inside
		klog.Infof("data: %s \nresponse length: %d\n", string(data1), len(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			klog.Error(err)
			return
		}
	}
}
