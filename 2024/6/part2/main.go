package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

func getRightCoords(pos [2]int, dir int) [2]int {
	switch dir {
	case UP:
		return [2]int{pos[0], pos[1] + 1}
	case RIGHT:
		return [2]int{pos[0] + 1, pos[1]}
	case DOWN:
		return [2]int{pos[0], pos[1] - 1}
	default:
		return [2]int{pos[0] - 1, pos[1]}
	}
}

func getRightVec(dir int) [2]int {
	switch dir {
	case UP:
		return [2]int{0, 1}
	case RIGHT:
		return [2]int{1, 0}
	case DOWN:
		return [2]int{0, -1}
	default:
		return [2]int{-1, 0}
	}
}

func rightDirHasObstacle(pos [2]int, dir int, guardMap map[[2]int]int) bool {
	rightVec := getRightVec(dir)
	testVec := [2]int{pos[0] + rightVec[0], pos[1] + rightVec[1]}
	for {
		testPos, ok := guardMap[testVec]
		if !ok {
			break
		}
		if testPos == OBSTACLE {
			return true
		}
		testVec = [2]int{testVec[0] + rightVec[0], testVec[1] + rightVec[1]}
	}

	return false
}

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

	obsPositions := 0
	obsCoords := [][2]int{}
	// move == 1, turn == 2
	previousMove := -1
	for {

		// determine next position
		guardVec := getNextVector(guardDirection)
		nextPos := [2]int{guardPosition[0] + guardVec[0], guardPosition[1] + guardVec[1]}

		nextElement, ok := guardMap[nextPos]
		if !ok {
			//we out of bounds
			break
		}

		// check. if the one to the right of us has already been visited.
		// if it has, then the NEXT position would be a good spot for an obstacle
		rightCoords := getRightCoords(guardPosition, guardDirection)
		// is it in bounds + has it been visited?
		rightElement, ok := guardMap[rightCoords]
		// does the right dir eventually have an obstacle?
		hasRightObstacle := ok && rightElement == VISITED && nextElement != OBSTACLE && previousMove == 1 && rightDirHasObstacle(guardPosition, guardDirection, guardMap)
		fmt.Printf("Guard Pos: %+v, NextPos: %+v,  Has Right Obstacle? %t\n", guardPosition, nextPos, hasRightObstacle)

		// spew.Dump(guardPosition, ok && rightElement == VISITED && nextElement != OBSTACLE && previousMove == 1)
		if hasRightObstacle {
			if !slices.Contains(obsCoords, nextPos) {
				obsCoords = append(obsCoords, nextPos)
				obsPositions += 1
			}

		}

		//

		// am i supposed to turn
		// yes - turn
		if nextElement == OBSTACLE {
			guardDirection = (guardDirection + 1) % 4
			previousMove = 2
		} else {
			// move
			// update previous + next
			guardMap[guardPosition] = VISITED
			guardMap[nextPos] = GUARD_VISITED

			// and move
			guardPosition = nextPos
			previousMove = 1
		}

	}
	numVisited := 0
	for _, val := range guardMap {
		if val == VISITED || val == GUARD_VISITED {
			numVisited += 1
		}
	}
	spew.Dump(obsCoords)
	spew.Dump(numVisited, obsPositions)

}
