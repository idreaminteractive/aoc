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
			if inputData[row][col] == "A" {
				// test it.
				startPoint := [2]int{row, col}
				if (getNorthEast(startPoint, inputData) == "M" && getSouthWest(startPoint, inputData) == "S") || (getNorthEast(startPoint, inputData) == "S" && getSouthWest(startPoint, inputData) == "M") {
					if (getNorthEast(startPoint, inputData) == "M" && getSouthWest(startPoint, inputData) == "S") || (getNorthEast(startPoint, inputData) == "S" && getSouthWest(startPoint, inputData) == "M") {
						// left off here.
						numFound += 1
					}

				}

			}
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

func hasXMAS(startPoint [2]int, input [][]string) bool {

	return false
}

func getNorthWest(startPoint [2]int, input [][]string) string {
	var point [2]int = [2]int{startPoint[0] - 1, startPoint[1] + 1}
	if isCoordInBounds(point, input) {
		return input[point[0]][point[1]]
	}
	return ""
}

func getNorthEast(startPoint [2]int, input [][]string) string {

	var point [2]int = [2]int{startPoint[0] - 1, startPoint[1] - 1}
	if isCoordInBounds(point, input) {
		return input[point[0]][point[1]]
	}
	return ""
}

func getSouthEast(startPoint [2]int, input [][]string) string {

	var point [2]int = [2]int{startPoint[0] + 1, startPoint[1] - 1}
	if isCoordInBounds(point, input) {
		return input[point[0]][point[1]]
	}
	return ""
}

func getSouthWest(startPoint [2]int, input [][]string) string {

	var point [2]int = [2]int{startPoint[0] + 1, startPoint[1] + 1}
	if isCoordInBounds(point, input) {
		return input[point[0]][point[1]]
	}
	return ""
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
