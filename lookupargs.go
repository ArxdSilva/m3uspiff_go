package main

import (
	"fmt"
	"github.com/dhowden/tag"
	"os"
	"strconv"
)

func lookupargs(line string) {
	file, err := os.Open(line)
	if err != nil {
		//fmt.Println(err)
		return
	}
	tagf, err := tag.ReadFrom(file)
	if err != nil {
		//fmt.Println(err)
		return
	}
	creator := tagf.Artist()
	album := tagf.Album()
	title := tagf.Title()
	trackNum, _ := tagf.Track()
	tags := []string{creator, album, title, strconv.Itoa(trackNum)}
	fmt.Println(tags)
}
