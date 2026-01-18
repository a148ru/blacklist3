package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"os"
)

func md5sum(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func loadMD5(path string) map[string]string {
	b, err := os.ReadFile(path)
	if err != nil {
		return map[string]string{}
	}
	var m map[string]string
	json.Unmarshal(b, &m)
	return m
}

func saveMD5(path string, data map[string]string) {
	b, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(path, b, 0644)
}
