package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal("Missing file input")
	}

	f, err := os.ReadFile(argsWithoutProg[0])
	if err != nil {
		log.Fatal(err)
	}
	data := string(f)

	// create our finite state machine
	var state rune = 'm'

	for _, r := range data {
		if r == state {
			spew.Dump("FOUND", r)
		}
	}

}
