package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X, Y int
}

func main() {

	day := "2024/day10"

	part1ex := part1(parseInput(day + "/input-ex.txt"))
	part1exExpected := 36
	log.Printf("part1 (example): %d\n", part1ex)
	if part1ex != part1exExpected {
		log.Fatalf("expecting %d\n", part1exExpected)
	}

	part1puzzle := part1(parseInput(day + "/input.txt"))
	log.Printf("part1 (puzzle): %d\n", part1puzzle)

	part2ex := part2(parseInput(day + "/input-ex.txt"))
	part2exExpected := 81
	log.Printf("part2 (example): %d\n", part2ex)
	if part2ex != part2exExpected {
		log.Fatalf("expecting %d\n", part2exExpected)
	}

	part2puzzle := part2(parseInput(day + "/input.txt"))
	log.Printf("part2 (puzzle): %d\n", part2puzzle)
}

func parseInput(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := make([]string, 0)
	fileBuf := bufio.NewReader(file)
	for {
		line, err := fileBuf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		input = append(input, strings.TrimSpace(line))
	}
	return input
}

func part1(input []string) int {
	topo := inputToTopo(input)
	score := 0
	for y, _ := range topo {
		for x, _ := range topo[0] {
			score += trailheadScore(topo, Position{x, y}, 0, 9, make(map[Position]bool))
		}
	}
	return score
}

func part2(input []string) int {
	topo := inputToTopo(input)
	score := 0
	for y, _ := range topo {
		for x, _ := range topo[0] {
			score += trailheadScore(topo, Position{x, y}, 0, 9, nil)
		}
	}
	return score
}

func inputToTopo(input []string) [][]int {
	topo := make([][]int, len(input))
	for y, _ := range input {
		topo[y] = make([]int, len(input[y]))
		for x, _ := range input[y] {
			n, err := strconv.Atoi(string(input[y][x]))
			if err != nil {
				log.Fatal(err)
			}
			topo[y][x] = n
		}
	}
	return topo
}

func trailheadScore(topo [][]int, position Position, expect int, end int, ends map[Position]bool) int {
	if position.X < 0 || position.Y < 0 || position.X >= len(topo[0]) || position.Y >= len(topo) {
		return 0
	}
	if topo[position.Y][position.X] != expect {
		return 0
	}
	if topo[position.Y][position.X] == end {
		if ends != nil {
			if _, exists := ends[position]; exists {
				return 0
			}
			ends[position] = true
		}
		return 1
	}
	score := 0
	score += trailheadScore(topo, Position{position.X - 1, position.Y}, expect+1, end, ends)
	score += trailheadScore(topo, Position{position.X + 1, position.Y}, expect+1, end, ends)
	score += trailheadScore(topo, Position{position.X, position.Y - 1}, expect+1, end, ends)
	score += trailheadScore(topo, Position{position.X, position.Y + 1}, expect+1, end, ends)
	return score
}
