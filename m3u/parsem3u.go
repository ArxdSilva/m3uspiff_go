package m3u

import (
	"bufio"
	"os"
	"strings"
)

var Lines = make([]string, 0)

func Parsem3u(file *os.File) {
	read := bufio.NewScanner(file)
	for read.Scan() {
		line := read.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		Lines = append(Lines, line)
	}
	return
}
