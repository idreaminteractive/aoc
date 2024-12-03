package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func absit(x int) int {
	if x < 0 {
		return x * -1
	}
	return x

}
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
	lines := strings.Split(data, "\n")

	numValid := 0
	for _, l := range lines {
		dir := 0
		fields := strings.Fields(l)

		lastVal, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		for i := 1; i < len(fields); i++ {
			val, err := strconv.Atoi(fields[i])
			// spew.Dump("vals", lastVal, val)
			if err != nil {
				log.Fatal(err)
			}

			abs := absit(lastVal - val)
			// invalid value
			if abs == 0 || abs > 3 {
				// spew.Dump("invalid, too far", lastVal, val, abs)
				break
			}

			// check direction
			// descending
			if lastVal > val {
				// last dir was ascending
				if dir > 0 {
					// spew.Dump("INVALID was desc")
					break
				} else if dir == 0 {
					dir = -1
				}
			} else {
				// we're ascending
				if dir < 0 {
					// spew.Dump("INVALID was asc")
					break
				} else if dir == 0 {
					dir = 1
				}
			}
			// we're at the end and we have not broke
			if i+1 == len(fields) {
				// spew.Dump("VALID")
				numValid += 1
			} else {
				// goto the next
				lastVal = val
			}

		}

	}
	spew.Dump(numValid)
}
