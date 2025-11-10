package client

import (
	"crypto/tls"
	"net/http"
)

func getNewClient(skipCertVerify bool) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipCertVerify},
	}

	return &http.Client{Transport: tr}
}
