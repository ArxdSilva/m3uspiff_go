package main

import (
	"fmt"
	"os"
)

func main() {
	command()
	if err == 1 {
		fmt.Println("m3uspiff takes exactly ONE argument")
		fmt.Println(err)
		os.Exit(1)
	}
	//parsem3u()
	//lookupargs()
	//makexml()
	print()
}
