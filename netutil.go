package netutil

import (
	"errors"
	"net"
	"os"
)

func GetLocalIpv4() (net.IP, error) {
	var ipv4 net.IP
	if hostname, err := os.Hostname(); err != nil {
		return ipv4, err
	} else {
		if ips, err := net.LookupIP(hostname); err != nil {
			return ipv4, err
		} else {
			for _, ip := range ips {
				if ipv4 = ip.To4(); ipv4 != nil {
					return ipv4, nil
				}
			}
			return ipv4, errors.New("failed to get local IP address")
		}
	}
}
