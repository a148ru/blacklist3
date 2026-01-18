package main

import (
	"io"
	"net/http"
	"os"
)

func loadSource(src Source) ([]byte, error) {
	if src.Type == "url" {
		resp, err := http.Get(src.Path)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return io.ReadAll(resp.Body)
	}

	return os.ReadFile(src.Path)
}
