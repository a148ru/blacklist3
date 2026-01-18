package main

import (
	"log"
	"os"
)

var logger *log.Logger

func initLogger() {
	logger = log.New(os.Stdout, "[md5loader] ", log.LstdFlags|log.Lshortfile)
}
