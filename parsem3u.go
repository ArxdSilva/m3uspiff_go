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
		if line != "" {
			if !strings.HasPrefix(line, "#") {
				lines = append(lines, line)
			}
		}
	}
	return
}
