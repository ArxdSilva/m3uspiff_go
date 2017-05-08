package main

import (
	"errors"
	"fmt"
	"github.com/ibrokemypie/m3uspiff_go/m3u"
	"github.com/ibrokemypie/m3uspiff_go/xspf"
	"os"
)

var bigmap = make(map[string]map[string]string)

func main() {

	if len(os.Args) != 2 {
		err := errors.New("m3uspiff takes exactly ONE argument")
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err == nil {
		m3u.Parsem3u(file)

		for _, line := range m3u.Lines {
			m3u.Lookupargs(line, bigmap)
		}
	}

	for _, tags := range bigmap {
		xspf.Makexml(tags)
	}

	fmt.Println(xspf.Output.String())
}
