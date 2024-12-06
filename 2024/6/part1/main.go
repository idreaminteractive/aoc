package main

import (
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

const EMPTY = 0
const GUARD = 1
const OBSTACLE = 2
const VISITED = 3

// If it's already been visited, but guard is currently there
const GUARD_VISITED = 4

const UP = 0
const RIGHT = 1
const DOWN = 2
const LEFT = 3

func getNextVector(direction int) [2]int {
	switch direction {
	case UP:
		return [2]int{-1, 0}
	case RIGHT:
		return [2]int{0, 1}
	case DOWN:
		return [2]int{1, 0}
	default:
		return [2]int{0, -1}
	}

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
	// each cell can have 4 states
	// guard, empty, obstacle, visted
	guardMap := map[[2]int]int{}
	guardPosition := [2]int{}
	guardDirection := UP
	lines := strings.Split(data, "\n")
	for row, l := range lines {
		for col, c := range l {
			switch c {
			case '^':
				// technically, it's both.
				guardPosition = [2]int{row, col}
				guardMap[[2]int{row, col}] = GUARD_VISITED
			case '#':
				guardMap[[2]int{row, col}] = OBSTACLE
			default:
				guardMap[[2]int{row, col}] = EMPTY
			}
		}
	}

	for {

		// determine next position
		guardVec := getNextVector(guardDirection)
		nextPos := [2]int{guardPosition[0] + guardVec[0], guardPosition[1] + guardVec[1]}

		element, ok := guardMap[nextPos]
		if !ok {
			//we out of bounds
			break
		}

		//

		// am i supposed to turn
		// yes - turn
		if element == OBSTACLE {
			guardDirection = (guardDirection + 1) % 4
		} else {
			// move
			// update previous + next
			guardMap[guardPosition] = VISITED
			guardMap[nextPos] = GUARD_VISITED

			// and move
			guardPosition = nextPos
		}

	}
	numVisited := 0
	for _, val := range guardMap {
		if val == VISITED || val == GUARD_VISITED {
			numVisited += 1
		}
	}
	spew.Dump(numVisited)

}
