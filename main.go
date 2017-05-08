package main

import (
	"fmt"
	"os"
)

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

	if err != nil {
		fmt.Println(err)
	}

	for _, tags := range tags {
		makexml(tags)
	}

	//print()
}

func errorh(err error) {
	fmt.Println(err)
	os.Exit(1)
}
