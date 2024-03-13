package utils

import "net"

func IsInternalIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	ipv4 := parsedIP.To4()
	if ipv4 == nil {
		return false
	}

	// 内网IP地址范围
	privateIPBlocks := []*net.IPNet{
		parseCIDR("10.0.0.0/8"),
		parseCIDR("172.16.0.0/12"),
		parseCIDR("192.168.0.0/16"),
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ipv4) {
			return true
		}
	}

	return false
}

func parseCIDR(cidr string) *net.IPNet {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		panic(err)
	}
	return ipNet
}

// 快速排序
