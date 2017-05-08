package xspf

import (
	"encoding/xml"
	"fmt"
)

func Makexml(tags map[string]string) {
	type track struct {
		Location string `xml:"location"`
		Title    string `xml:"title,omitempty"`
		TrackNum string `xml:"trackNum,omitempty"`
		Creator  string `xml:"creator,omitempty"`
		Album    string `xml:"album,omitempty"`
	}

	v := &track{
		Location: tags["location"],
		Title:    tags["title"],
		TrackNum: tags["trackNum"],
		Creator:  tags["creator"],
		Album:    tags["album"],
	}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(string(output))
}

//func Makexml(tags map[string]string) {
//fmt.Println(tags)
//xml.Marshaler
//output, err := xml.MarshalIndent(data, "", "  ")
//}
