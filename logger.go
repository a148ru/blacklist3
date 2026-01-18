package main

import (
	"log"
	"os"
)

var logger *log.Logger

func initLogger() {
	logger = log.New(os.Stdout, "[blacklist] ", log.LstdFlags|log.Lshortfile)
}
