package m3u

import (
	"github.com/dhowden/tag"
	"os"
	"strconv"
)

var entries = 0

func Lookupargs(line string, err error, bigmap map[string]map[string]string) {
	tags := make(map[string]string)
	tags["entry"] = strconv.Itoa(entries)
	tags["location"] = line

	file, err1 := os.Open(line)
	tagf, err2 := tag.ReadFrom(file)
	if err1 != nil || err2 != nil {
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
		end(tags, bigmap)
		return
	}

	artist := tagf.Artist()
	album := tagf.Album()
	title := tagf.Title()
	track, _ := tagf.Track()

	tags["creator"] = artist
	tags["album"] = album
	tags["title"] = title
	tags["trackNum"] = strconv.Itoa(track)

	end(tags, bigmap)
}

func end(tags map[string]string, bigmap map[string]map[string]string) {
	bigmap[strconv.Itoa(entries)] = tags
	entries++
	return
}
