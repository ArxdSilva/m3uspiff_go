package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	argcheck()
	fmt.Println(os.Args[1])
}

func argcheck() {
	if len(os.Args) != 2 {
		err := errors.New("m3uspiff takes exactly ONE argument \n")
		fmt.Print(err)
		os.Exit(1)
	}
}
