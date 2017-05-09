package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/ibrokemypie/m3uspiff_go/m3u"
	"github.com/ibrokemypie/m3uspiff_go/xspf"
)

var trackMap = map[string]*m3u.Track{}
var entries = 0

func main() {
	if len(os.Args) != 2 {
		err := errors.New("m3uspiff takes exactly ONE argument")
		fmt.Println(err)
		os.Exit(1)
	}

	file, _ := os.Open(os.Args[1])
	m3u.Parsem3u(file)
	for _, line := range m3u.Lines {
		entry, err := m3u.Lookupargs(line)
		if err != nil {
			panic(err)
		}
		trackMap[strconv.Itoa(entries)] = entry
		entries++
	}

	err := xspf.Makexml(trackMap)
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}
