package xspf

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/ibrokemypie/m3uspiff_go/m3u"
)

type playlist struct {
	Tracklist map[string]*m3u.Track `xml:"tracklist"`
	Version   int                   `xml:"version,attr"`
	Xmlns     string                `xml:"xmlns,attr"`
}

var Output bytes.Buffer

func Makexml(trackMap map[string]*m3u.Track) error {
	Output.WriteString(xml.Header)
	final := &playlist{
		Version:   1,
		Xmlns:     "http://xspf.org/ns/0/",
		Tracklist: trackMap,
	}

	oops, err := xml.MarshalIndent(final, "  ", "    ")
	if err != nil {
		return err
	}

	Output.Write(oops)
	fmt.Println(Output.String())
	return err
}
