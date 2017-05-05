package main

import (
	"bufio"
	"os"
	"strings"
)

var lines = make([]string, 0)

func parsem3u(file *os.File) {
	read := bufio.NewScanner(file)
	for read.Scan() {
		line := read.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		lines = append(lines, line)
	}
	return
}
