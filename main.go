package main

import (
	"errors"
	"fmt"
	"github.com/ibrokemypie/m3uspiff_go/m3u"
	"github.com/ibrokemypie/m3uspiff_go/xspf"
	"os"
)

var err error
var file *os.File
var bigmap = make(map[string]map[string]string)

func command() {
	if len(os.Args) != 2 {
		err = errors.New("m3uspiff takes exactly ONE argument")
		return
	}
	file, err = os.Open(os.Args[1])
}

func main() {

	command()

	if err != nil {
		errorh(err)
	}

	m3u.Parsem3u(file)

	if err != nil {
		errorh(err)
	}

	for _, line := range m3u.Lines {
		m3u.Lookupargs(line, err, bigmap)
	}

	for _, tags := range bigmap {
		xspf.Makexml(tags)
	}
}

func errorh(err error) {
	fmt.Println(err)
	os.Exit(1)
}
