package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

const (
	CELL_EMPTY = '.'
	CELL_OBSTR = '#'
	CELL_GUARD = '^'
	CELL_VISIT = 'X'
	CELL_OUTSI = '!'
	CELL_LOOOP = 'L'
)

type Direction struct {
	X, Y int
}

func (d Direction) TurnRight() Direction {
	switch d {
	case DIR_UP:
		return DIR_RIGHT
	case DIR_RIGHT:
		return DIR_DOWN
	case DIR_DOWN:
		return DIR_LEFT
	case DIR_LEFT:
		return DIR_UP
	}
	log.Fatal("invalid direction")
	return Direction{}
}

var (
	DIR_UP    = Direction{X: 0, Y: -1}
	DIR_RIGHT = Direction{X: 1, Y: 0}
	DIR_DOWN  = Direction{X: 0, Y: 1}
	DIR_LEFT  = Direction{X: -1, Y: 0}
)

type Position struct {
	X, Y int
	D    Direction
}

func (p Position) Ahead() Position {
	return Position{X: p.X + p.D.X, Y: p.Y + p.D.Y, D: p.D}
}

func main() {

	part1ex := part1(parsePlan("day6/input-ex.txt"))
	log.Printf("distinct positions (example): %d\n", part1ex)
	if part1ex != 41 {
		log.Fatalf("expecting %d distinct positions\n", 41)
	}

	part1puzzle := part1(parsePlan("day6/input.txt"))
	log.Printf("distinct positions (puzzle): %d\n", part1puzzle)

	part2ex := part2(parsePlan("day6/input-ex.txt"))
	log.Printf("possible loops (example): %d\n", part2ex)
	if part2ex != 6 {
		log.Fatalf("expecting %d possible loops\n", 6)
	}

	part2puzzle := part2(parsePlan("day6/input.txt"))
	log.Printf("possible loops (puzzle): %d\n", part2puzzle)
}

func parsePlan(fileName string) [][]rune {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileBuf := bufio.NewReader(file)
	plan := make([][]rune, 0)
	for {
		line, err := fileBuf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		plan = append(plan, []rune(line))
	}
	return plan
}

func part1(plan [][]rune) int {
	guardWalkUntilOutsideOrInLoop(plan)
	return numberOfVisitedCell(plan)
}

func part2(plan [][]rune) int {
	return numberOfPossibleLoop(plan)
}

func numberOfPossibleLoop(plan [][]rune) int {
	loops := 0
	for y, cells := range plan {
		for x, cell := range cells {
			if cell == CELL_EMPTY {
				nplan := copyPlan(plan)
				nplan[y][x] = CELL_OBSTR
				if isGuardWalkInLoop(nplan) {
					loops++
				}
			}
		}
	}
	return loops
}

func isGuardWalkInLoop(plan [][]rune) bool {
	cell := guardWalkUntilOutsideOrInLoop(plan)
	if cell == CELL_LOOOP {
		return true
	}
	return false
}

func guardWalkUntilOutsideOrInLoop(plan [][]rune) rune {
	visitedPositions := make([]Position, 0)
	position := findGuardPosition(plan)
	for {
		cell, ahead := lookAhead(plan, position)
		switch cell {
		case CELL_OUTSI:
			fallthrough
		case CELL_EMPTY:
			fallthrough
		case CELL_VISIT:
			plan[position.Y][position.X] = CELL_VISIT
			visitedPositions = append(visitedPositions, position)
			position = ahead
		case CELL_OBSTR:
			position.D = position.D.TurnRight()
		}
		if cell == CELL_OUTSI {
			return cell
		}
		for _, visitedPosition := range visitedPositions {
			if position == visitedPosition {
				return CELL_LOOOP
			}
		}
	}
}

func findGuardPosition(plan [][]rune) Position {
	for y, cells := range plan {
		for x, cell := range cells {
			if cell == CELL_GUARD {
				return Position{x, y, DIR_UP}
			}
		}
	}
	log.Fatal("no guard found")
	return Position{}
}

func numberOfVisitedCell(plan [][]rune) int {
	visited := 0
	for _, cells := range plan {
		for _, cell := range cells {
			if cell == CELL_VISIT {
				visited++
			}
		}
	}
	return visited
}

func copyPlan(plan [][]rune) [][]rune {
	nplan := make([][]rune, len(plan))
	for y, cells := range plan {
		nplan[y] = make([]rune, len(cells))
		for x, cell := range cells {
			nplan[y][x] = cell
		}
	}
	return nplan
}

func lookAhead(plan [][]rune, position Position) (rune, Position) {
	ahead := position.Ahead()
	if ahead.X < 0 || ahead.Y < 0 || ahead.X >= len(plan[0]) || ahead.Y >= len(plan) {
		return CELL_OUTSI, ahead
	}
	return plan[ahead.Y][ahead.X], ahead
}
