package main

import (
	"log"
	"os"
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

	on := true
	i := 0
	for {

		if on {
			// we're looking for mul( x, y )
		}

	}

}
