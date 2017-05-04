package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("m3uspiff takes exactly ONE argument \n")
		os.Exit(1)
	}
	print()
}
