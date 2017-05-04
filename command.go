package main

import (
	"os"
)

var err = 0

func command() {
	if len(os.Args) != 2 {
		err = 1
	}
}
