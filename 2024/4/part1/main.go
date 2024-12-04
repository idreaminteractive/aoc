package main

import (
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// we one shot it, let's fucking go
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

	// read letters into a matrix [][]
	// know x + know y
	lines := strings.Split(data, "\n")
	// spew.Dump(lines)

	inputData := make([][]string, len(lines))
	for y, l := range lines {
		inputData[y] = make([]string, len(l))
		for x, c := range l {

			inputData[y][x] = string(c)
		}
	}
	// ok - we're good. now, go through and find xmas.
	// state := ""
	numFound := 0
	for row := 0; row < len(inputData); row++ {
		for col := 0; col < len(inputData[row]); col++ {
			// fmt.Printf("%s ", inputData[row][col])
			// so, in here, i want to search my eight patterns
			if hasNorth([2]int{row, col}, inputData) {
				numFound += 1
			}
			if hasNorthWest([2]int{row, col}, inputData) {
				numFound += 1
			}
			if hasWest([2]int{row, col}, inputData) {
				numFound += 1
			}
			if hasSouthWest([2]int{row, col}, inputData) {
				numFound += 1
			}
			if hasSouth([2]int{row, col}, inputData) {
				numFound += 1
			}
			if hasSouthEast([2]int{row, col}, inputData) {
				numFound += 1
			}
			if hasEast([2]int{row, col}, inputData) {
				numFound += 1
			}
			if hasNorthEast([2]int{row, col}, inputData) {
				numFound += 1
			}
		}
	}
	spew.Dump(numFound)

}

func isCoordInBounds(coord [2]int, input [][]string) bool {
	if coord[0] >= len(input) {
		return false
	}
	if coord[0] < 0 {
		return false
	}
	if coord[1] >= len(input[coord[0]]) {
		return false
	}
	if coord[1] < 0 {
		return false
	}
	return true
}

func testPattern(startPoint [2]int, pattern [4][2]int, input [][]string) bool {
	state := ""
	for _, p := range pattern {
		testCoord := [2]int{startPoint[0] + p[0], startPoint[1] + p[1]}
		if isCoordInBounds(testCoord, input) {
			// ok - it's legit
			state += input[testCoord[0]][testCoord[1]]
		}
	}
	return state == "XMAS"
}

func hasNorth(startPoint [2]int, input [][]string) bool {
	// bounds check each thing

	pattern := [4][2]int{{0, 0}, {-1, 0}, {-2, 0}, {-3, 0}}
	return testPattern(startPoint, pattern, input)
}

func hasNorthWest(startPoint [2]int, input [][]string) bool {
	// bounds check each thing

	pattern := [4][2]int{{0, 0}, {-1, 1}, {-2, 2}, {-3, 3}}
	return testPattern(startPoint, pattern, input)
}

func hasWest(startPoint [2]int, input [][]string) bool {
	// bounds check each thing

	pattern := [4][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	return testPattern(startPoint, pattern, input)
}

func hasSouthWest(startPoint [2]int, input [][]string) bool {
	// bounds check each thing

	pattern := [4][2]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}}
	return testPattern(startPoint, pattern, input)
}

func hasSouth(startPoint [2]int, input [][]string) bool {
	// bounds check each thing

	pattern := [4][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
	return testPattern(startPoint, pattern, input)
}

func hasSouthEast(startPoint [2]int, input [][]string) bool {
	// bounds check each thing

	pattern := [4][2]int{{0, 0}, {1, -1}, {2, -2}, {3, -3}}
	return testPattern(startPoint, pattern, input)
}

func hasEast(startPoint [2]int, input [][]string) bool {
	// bounds check each thing

	pattern := [4][2]int{{0, 0}, {0, -1}, {0, -2}, {0, -3}}
	return testPattern(startPoint, pattern, input)
}

func hasNorthEast(startPoint [2]int, input [][]string) bool {
	// bounds check each thing

	pattern := [4][2]int{{0, 0}, {-1, -1}, {-2, -2}, {-3, -3}}
	return testPattern(startPoint, pattern, input)
}
