// Copyright 2018 The QOS Authors

package utils

import "net"

func NewLocalIPNet() []*net.IPNet {
	var ipNets []*net.IPNet
	_, ipNet1, _ := net.ParseCIDR("127.0.0.1")
	ipNets = append(ipNets, ipNet1)

	_, ipNet2, _ := net.ParseCIDR("192.168.1.0/24")
	ipNets = append(ipNets, ipNet2)

	return ipNets
}
