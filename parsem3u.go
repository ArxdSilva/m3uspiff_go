package main

import (
	"bufio"
	"fmt"
	"os"
)

func parsem3u(file *os.File) {
	read := bufio.NewScanner(file)
	for read.Scan() {
		fmt.Println(read.Text())
	}
}
