package main

import (
	"io"
	"net/http"
	"os"
	"time"
	"fmt"
	"crypto/tls"
)

var httpClient *http.Client
/*
func initHTTP(timeoutSec int) {
	httpClient = &http.Client{
		Timeout: time.Duration(timeoutSec) * time.Second,
	}
}
*/

func initHTTP(cfg HTTPConf) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: cfg.InsecureSkipVerify,
		},
	}

	httpClient = &http.Client{
		Timeout:   time.Duration(cfg.TimeoutSeconds) * time.Second,
		Transport: tr,
	}
}

func loadSource(src Source) ([]byte, error) {
	if src.Type == "url" {
		resp, err := httpClient.Get(src.Path)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
		}

		return io.ReadAll(resp.Body)
	}

	return os.ReadFile(src.Path)
}
