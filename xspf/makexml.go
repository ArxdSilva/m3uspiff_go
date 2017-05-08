package xspf

import (
	"bytes"
	"encoding/xml"
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
	newoutput, _ := xml.MarshalIndent(values, "  ", "    ")
	Output.Write(newoutput)
	//Output = (Output + newoutput)

	//(Output, _ = xml.MarshalIndent(values, "  ", "    "))

	return
	//fmt.Println(string(output))
}
