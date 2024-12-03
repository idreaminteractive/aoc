package main

import (
	"log"
	"os"
	"regexp"
	"strconv"

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
	spew.Dump(data)
	var re = regexp.MustCompile(`(?m)(mul\((\d{1,3}),(\d{1,3})\))`)

	matches := re.FindAllStringSubmatch(data, -1)

	val := 0
	for _, m := range matches {
		if len(m) == 4 {
			m1, _ := strconv.Atoi(m[2])
			m2, _ := strconv.Atoi(m[3])
			val += m1 * m2

		}
	}
	spew.Dump(val)
	// for i, match := range re.FindAllString(data, -1) {
	// 	fmt.Println(match, "found at index", i)
	// }
}
