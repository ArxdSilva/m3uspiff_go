package m3u

import (
	"os"
	"strconv"

	"github.com/dhowden/tag"
)

type Track struct {
	Creator  string `xml:"creator,omitempty"`
	Album    string `xml:"album,omitempty"`
	Title    string `xml:"title,omitempty"`
	TrackNum string `xml:"trackNum,omitempty"`
	Location string `xml:"location"`
	Entry    string
}

var entries = 0

func Lookupargs(line string) (*Track, error) {
	file, err := os.Open(line)
	if err != nil {
		return nil, err
	}
	tagf, err := tag.ReadFrom(file)
	if err != nil {
		return nil, err
	}
	track, _ := tagf.Track()
	entry := Track{
		Creator:  tagf.Artist(),
		Album:    tagf.Album(),
		Title:    tagf.Title(),
		TrackNum: strconv.Itoa(track),
		Entry:    strconv.Itoa(entries),
		Location: line,
	}
	entries++
	return &entry, err
}

func end(tags map[string]string, bigmap map[string]map[string]string) {
	bigmap[strconv.Itoa(entries)] = tags

	return
}
