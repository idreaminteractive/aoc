package main

import (
	"log"
	"os"
	"strconv"
	"strings"

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

	left := []int{}
	right := []int{}
	lines := strings.Split(data, "\n")

	for _, l := range lines {
		splitData := strings.Split(l, " ")
		leftNum, err := strconv.Atoi(splitData[0])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, int(leftNum))

		rightNum, err := strconv.Atoi(splitData[len(splitData)-1])
		if err != nil {
			log.Fatal(err)
		}
		right = append(right, int(rightNum))
	}

	scores := 0
	for i := range left {
		numFound := 0

		for j := range right {
			if left[i] == right[j] {
				numFound += 1
			}
		}
		scores += numFound * left[i]
	}
	spew.Dump(scores)

}
