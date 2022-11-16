package common

import "net/http"

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
