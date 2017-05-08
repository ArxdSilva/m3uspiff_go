package main

import (
	"fmt"
	"os"
)

func main() {

	command()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	parsem3u(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, line := range lines {
		lookupargs(line)
	}

	makexml()
	//print()
}
