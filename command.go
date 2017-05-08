package main

import (
	"errors"
	"os"
)

var err error
var file *os.File

func command() {
	if len(os.Args) != 2 {
		err = errors.New("m3uspiff takes exactly ONE argument")
		return
	}
	file, err = os.Open(os.Args[1])
}
