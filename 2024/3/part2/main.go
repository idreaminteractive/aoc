package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

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
	var re = regexp.MustCompile(`(?m)(mul\(\d{1,3},\d{1,3}\))`)
	var str = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}
}
