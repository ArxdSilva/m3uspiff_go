package xspf

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
)

var Output bytes.Buffer

func Makexml(tags map[string]string) {
	type track struct {
		Location string `xml:"location"`
		Title    string `xml:"title,omitempty"`
		TrackNum string `xml:"trackNum,omitempty"`
		Creator  string `xml:"creator,omitempty"`
		Album    string `xml:"album,omitempty"`
	}

	values := &track{
		Location: tags["location"],
		Title:    tags["title"],
		TrackNum: tags["trackNum"],
		Creator:  tags["creator"],
		Album:    tags["album"],
	}
	Output.WriteString("\n")
	newoutput, err := xml.MarshalIndent(values, "  ", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Output.Write(newoutput)

	return
}
