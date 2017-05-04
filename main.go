package main

import (
	"fmt"
	"os"
)

func main() {
	command()

	if err == 1 {
		fmt.Println("m3uspiff takes exactly ONE argument")
		os.Exit(1)
	}
	var file, err = os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	parsem3u(file)
	//lookupargs()
	//makexml()
	//print()
}
