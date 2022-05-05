package main

import (
	"testing"
)

// func antiChristina(realIP string) string {
// 	if valid.IsIP(realIP) {
// 		return realIP
// 	}
// 	return "Bazinga!"
// }

func TestAntiChristinaIPv4(t *testing.T) {
	ip := "192.168.0.1"
	ans := antiChristina(ip)
	if ans != ip {

		t.Errorf("antiChristina('%s') = %s; want %s", ip, ans, ip)
	}
}

func TestAntiChristinaIPv6(t *testing.T) {
	ip := "2001:4860:4860::8888"
	ans := antiChristina(ip)
	if ans != ip {

		t.Errorf("antiChristina('%s') = %s; want %s", ip, ans, ip)
	}
}

func TestAntiChristinaNotIP(t *testing.T) {
	ip := "corgibutt"
	ans := antiChristina(ip)
	if ans != "Bazinga!" {

		t.Errorf("antiChristina('%s') = %s; want %s", ip, ans, ip)
	}
}
