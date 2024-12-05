package main

import (
	"log"
	"os"
	"slices"
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

	pList := map[int][]int{}

	lines := strings.Split(data, "\n")
	isParsingInput := false
	midSum := 0
	for _, line := range lines {

		if line == "" {
			// we  have reached the parsing input
			isParsingInput = true
			continue
		}
		if !isParsingInput {
			s := strings.Split(line, "|")
			key, _ := strconv.Atoi(s[0])
			val, _ := strconv.Atoi(s[1])
			pList[key] = append(pList[key], val)
		} else {

			s := strings.Split(line, ",")
			iarr := []int{}
			isValid := true
		out:
			for _, i := range s {
				v, _ := strconv.Atoi(i)
				exList, ok := pList[v]
				if !ok {
					iarr = append(iarr, v)
				} else {
					// check and see if this is excluded against iarr
					for _, xl := range exList {
						if slices.Contains(iarr, xl) {
							isValid = false
							break out
						}
					}
					// if not foudn, it's good
					iarr = append(iarr, v)
				}
			}
			// find the mid

			if isValid {
				if len(iarr) > 0 {
					if len(iarr) == 1 {
						midSum += iarr[0]
					} else {
						midSum += iarr[len(iarr)/2]
					}
				}
			}

		}

	}
	spew.Dump(midSum)

}
