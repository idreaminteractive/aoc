package main

import (
	"log"
	"os"
	"regexp"
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
	// spew.Dump(data)
	var re = regexp.MustCompile(`(?m)((mul\((\d{1,3}),(\d{1,3})\))|(don't\(\))|(do\(\)))`)

	matches := re.FindAllStringSubmatch(data, -1)
	// spew.Dump(matches)
	val := 0

	do := true
	for _, m := range matches {
		if strings.HasPrefix(m[0], "mul") {
			if do {
				m1, _ := strconv.Atoi(m[3])
				m2, _ := strconv.Atoi(m[4])
				val += m1 * m2
			}
		}
		if strings.HasPrefix(m[0], "don't") {
			do = false
		}
		if strings.HasPrefix(m[0], "do()") {
			do = true
		}
	}
	spew.Dump(val)
	// for i, match := range re.FindAllString(data, -1) {
	// 	fmt.Println(match, "found at index", i)
	// }
}
