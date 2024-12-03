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

func fits(arr []int) bool {

	dir := 0
	for i := 1; i < len(arr); i++ {
		lastVal := arr[i-1]
		val := arr[i]
		abs := absit(lastVal - val)
		// invalid value
		if abs == 0 || abs > 3 {
			return false
		}

		if lastVal > val {
			// last dir was ascending
			if dir > 0 {
				// spew.Dump("INVALID was desc")
				return false
			} else if dir == 0 {
				dir = -1
			}
		} else {
			// we're ascending
			if dir < 0 {

				return false
			} else if dir == 0 {
				dir = 1
			}
		}
		// we're at the end and we have not broke
		// if i+1 == len(arr) {
		// 	return true
		// }
	}
	// made it past all the bad checks
	return true
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
		fields := strings.Fields(l)

		arr := make([]int, 0)
		for _, n := range fields {
			num, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			arr = append(arr, num)
		}

		c := append([]int(nil), arr...)

		index := 0
		for {
			if index == len(arr) {
				break
			}
			spew.Dump("test", c)
			if fits(c) {
				numValid += 1
				spew.Dump("SAFE", c)
				break
			} else {
				c = append(arr[:index], arr[index+1:]...)
				index += 1
			}
		}

	}
	spew.Dump("numvalid", numValid)
}
