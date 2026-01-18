package main

import "testing"

func TestMD5(t *testing.T) {
	data := []byte("test")
	sum := md5sum(data)

	if sum != "098f6bcd4621d373cade4e832627b4f6" {
		t.Fatal("Неверная MD5 сумма")
	}
}

func TestRegex(t *testing.T) {
	line := "block 192.168.1.0/24 now"
	if !netRegex.MatchString(line) {
		t.Fatal("IPv4 сеть не найдена")
	}
}
