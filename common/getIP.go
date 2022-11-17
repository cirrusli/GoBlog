package common

import (
	"net"
	"net/http"
	"strings"
)

func GetRemoteIP(r *http.Request) string {
	var remoteAddr string
	remoteAddr = r.RemoteAddr
	if remoteAddr != "" {
		return remoteAddr
	}
	//ipv4
	remoteAddr = r.Header.Get("ipv4")
	if remoteAddr != "" {
		return remoteAddr
	}
	remoteAddr = r.Header.Get("XForwardedFor")
	if remoteAddr != "" {
		return remoteAddr
	}
	// X-Forwarded-For
	remoteAddr = r.Header.Get("X-Forwarded-For")
	if remoteAddr != "" {
		return remoteAddr
	}
	// X-Real-Ip
	remoteAddr = r.Header.Get("X-Real-Ip")
	if remoteAddr != "" {
		return remoteAddr
	} else {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

// ClientPublicIP 尽最大努力实现获取客户端公网 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" && !HasLocalIPAddr(ip) {
			return ip
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" && !HasLocalIPAddr(ip) {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		if !HasLocalIPAddr(ip) {
			return ip
		}
	}

	return ""
}

// HasLocalIPAddr 检测 IP 地址字符串是否是内网地址
func HasLocalIPAddr(ip string) bool {
	return HasLocalIP(net.ParseIP(ip))
}

//内网ip
var localNetworks = []net.IP{
	[]byte("10.0.0.0/8"),
	[]byte("169.254.0.0/16"),
	[]byte("172.16.0.0/12"),
	[]byte("172.17.0.0/12"),
	[]byte("172.18.0.0/12"),
	[]byte("172.19.0.0/12"),
	[]byte("172.20.0.0/12"),
	[]byte("172.21.0.0/12"),
	[]byte("172.22.0.0/12"),
	[]byte("172.23.0.0/12"),
	[]byte("172.24.0.0/12"),
	[]byte("172.25.0.0/12"),
	[]byte("172.26.0.0/12"),
	[]byte("172.27.0.0/12"),
	[]byte("172.28.0.0/12"),
	[]byte("172.29.0.0/12"),
	[]byte("172.30.0.0/12"),
	[]byte("172.31.0.0/12"),
	[]byte("192.168.0.0/16"),
}

// HasLocalIP 检测 IP 地址是否是内网地址
func HasLocalIP(ip net.IP) bool {
	for _, network := range localNetworks {
		if network.Equal(ip) {
			return true
		}
	}

	return ip.IsLoopback()
}
