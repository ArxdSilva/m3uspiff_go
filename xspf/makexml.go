package xspf

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
)

type playlist struct {
	Tracklist *tracklist `xml:"tracklist"`
	Version   int        `xml:"version,attr"`
	Xmlns     string     `xml:"xmlns,attr"`
}

type tracklist struct {
	Track *track `xml:"track"`
}

type track struct {
	XMLName  xml.Name `xml:"track"`
	Location string   `xml:"location"`
	Title    string   `xml:"title,omitempty"`
	TrackNum string   `xml:"trackNum,omitempty"`
	Creator  string   `xml:"creator,omitempty"`
	Album    string   `xml:"album,omitempty"`
}

var Output bytes.Buffer
var value xml.CharData
var entries *track
var Tracklist *tracklist

func Makexml(tags map[string]map[string]string) {
	for _, entry := range tags {
		entries = &track{
			Location: entry["location"],
			Title:    entry["title"],
			TrackNum: entry["trackNum"],
			Creator:  entry["creator"],
			Album:    entry["album"],
		}
		//value, _ = xml.MarshalIndent(entries, " ", "  ")
		//Output.Write(value)
		Tracklist = &tracklist{
			Track: entries,
		}
	}

	Output.WriteString(xml.Header)

	final := &playlist{
		Version:   1,
		Xmlns:     "http://xspf.org/ns/0/",
		Tracklist: Tracklist,
	}

	oops, err := xml.MarshalIndent(final, "  ", "    ")
	Output.Write(oops)
	println(Output.String())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}
