package main

import (
	"fmt"
	"os"
)

func main() {
	argcheck()
	fmt.Println(os.Args[1])
}

func argcheck() {
	if len(os.Args) != 2 {
		fmt.Println("m3uspiff takes exactly ONE argument \n")
		os.Exit(1)
	}
}
