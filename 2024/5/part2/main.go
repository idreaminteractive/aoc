package main

import (
	"fmt"
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

			for _, i := range s {
				v, _ := strconv.Atoi(i)
				exList, ok := pList[v]
				if !ok {
					iarr = append(iarr, v)
				} else {
					// check and see if this is excluded against iarr
					toInsert := 1000
					fmt.Printf("Ex: %+v, v: %d, iarr: %+v\n", exList, v, iarr)
					// out:
					for _, xl := range exList {

						// left off here
						foundAt := slices.Index(iarr, xl)

						if toInsert > foundAt && foundAt > -1 {
							toInsert = foundAt
						}
					}
					if toInsert == 1000 {
						// if not foudn, it's good
						iarr = append(iarr, v)
					} else {
						isValid = false
						iarr = slices.Insert(iarr, toInsert, v)
					}

				}
			}
			// find the mid

			if !isValid {
				spew.Dump(iarr)
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
