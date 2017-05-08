package main

import (
	"fmt"
	"os"
)

var bigmap = make(map[string]map[string]string)

func main() {

	command()

	if err != nil {
		errorh(err)
	}

	parsem3u(file)

	if err != nil {
		errorh(err)
	}

	for _, line := range lines {
		lookupargs(line)
	}

	for _, tags := range bigmap {
		makexml(tags)
	}

	//print()
}

func errorh(err error) {
	fmt.Println(err)
	os.Exit(1)
}
