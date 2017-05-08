package main

import (
	"github.com/dhowden/tag"
	"os"
	"strconv"
)

var tags []string

func lookupargs(line string) {
	file, err1 := os.Open(line)
	tagf, err2 := tag.ReadFrom(file)
	if err1 != nil || err2 != nil {
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
		return
	}

	creator := tagf.Artist()
	album := tagf.Album()
	title := tagf.Title()
	trackNum, _ := tagf.Track()
	tags = []string{creator, album, title, strconv.Itoa(trackNum)}
	return
}
