package main

import (
	"crypto/tls"

	"github.com/labstack/gommon/log"
)

// Implement a basic JA3 fingerprint
func doJA3() {
	cs := tls.CipherSuites()
	log.Info(cs)
}
