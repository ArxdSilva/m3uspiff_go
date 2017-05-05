package main

import (
	"github.com/dhowden/tag"
	"log"
	"os"
)

func lookupargs(line string) {
	//println(line)
	file, err := os.Open(line)
	if err != nil {
		//log.Print(err)
		return
	}
	tags, err := tag.ReadFrom(file)
	if err != nil {
		log.Print(err)
		return
	}
	println(tags.Title())

}
