package main

import (
	"bufio"
	"os"
	"strings"
)

func parsem3u(file *os.File) {
	read := bufio.NewScanner(file)
	for read.Scan() {
		var line = read.Text()
		if line != "" {
			if !strings.HasPrefix(line, "#") {
				lookupargs(line)
			}
		}
	}
}
