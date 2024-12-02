package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func absit(x int64) int64 {
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
	left := []int64{}
	right := []int64{}
	lines := strings.Split(data, "\n")

	for _, l := range lines {
		splitData := strings.Split(l, " ")
		leftNum, err := strconv.Atoi(splitData[0])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, int64(leftNum))

		rightNum, err := strconv.Atoi(splitData[len(splitData)-1])
		if err != nil {
			log.Fatal(err)
		}
		right = append(right, int64(rightNum))
	}

	slices.Sort(left)
	slices.Sort(right)
	var distance int64 = 0
	for i := range left {
		distance += absit(left[i] - right[i])
	}
	spew.Dump(distance)
}
