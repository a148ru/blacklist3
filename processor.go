package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//var netRegex = regexp.MustCompile(`\b\d{1,3}(\.\d{1,3}){3}/\d{1,2}\b`)
// \b(?:(?:25[0-5]|2[0-4]\d|1?\d?\d)(?:\.(?:25[0-5]|2[0-4]\d|1?\d?\d)){3})(?:/(?:3[0-2]|[12]?\d))?\b

var netRegex = regexp.MustCompile(`\b(?:(?:25[0-5]|2[0-4]\d|1?\d?\d)(?:\.(?:25[0-5]|2[0-4]\d|1?\d?\d)){3})(?:/(?:3[0-2]|[12]?\d))?\b`)
func processData(data []byte, outDir, name string) error {
	lines := strings.Split(string(data), "\n")

	outPath := filepath.Join(outDir, name+".conf")
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range lines {
		if netRegex.MatchString(line) {
			network := netRegex.FindString(line)
			fmt.Fprintf(f, "route %s reject;\n", network)
		}
	}
	return nil
}
