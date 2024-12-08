package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Position struct {
	X, Y int
}

func main() {

	day := "day8"
	part1ex := part1(parseInput(day + "/input-ex.txt"))
	part1exExpected := 14
	log.Printf("part1 (example): %d\n", part1ex)
	if part1ex != part1exExpected {
		log.Fatalf("expecting %d\n", part1exExpected)
	}

	part1puzzle := part1(parseInput(day + "/input.txt"))
	log.Printf("part1 (puzzle): %d\n", part1puzzle)

	part2ex := part2(parseInput(day + "/input-ex.txt"))
	part2exExpected := 34
	log.Printf("part2 (example): %d\n", part2ex)
	if part2ex != part2exExpected {
		log.Fatalf("expecting %d\n", part2exExpected)
	}

	part2puzzle := part2(parseInput(day + "/input.txt"))
	log.Printf("part2 (puzzle): %d\n", part2puzzle)
}

func parseInput(fileName string) [][]rune {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := make([][]rune, 0)
	fileBuf := bufio.NewReader(file)
	for {
		line, err := fileBuf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		input = append(input, []rune(strings.TrimSpace(line)))
	}
	return input
}

func part1(antennasPlan [][]rune) int {
	antinodesPlan := makeAntinodesPlan(antennasPlan)
	populateAntinodesPlan(antennasPlan, antinodesPlan, false)
	return countUniqueAntinodes(antinodesPlan)
}

func part2(antennasPlan [][]rune) int {
	antinodesPlan := makeAntinodesPlan(antennasPlan)
	populateAntinodesPlan(antennasPlan, antinodesPlan, true)
	return countUniqueAntinodes(antinodesPlan)
}

func findAllFrequencies(antennasPlan [][]rune) map[rune][]Position {
	frequencies := make(map[rune][]Position)
	for y, _ := range antennasPlan {
		for x, _ := range antennasPlan[y] {
			f := antennasPlan[y][x]
			if f != '.' {
				if _, exists := frequencies[f]; !exists {
					frequencies[f] = make([]Position, 0)
				}
				frequencies[f] = append(frequencies[f], Position{x, y})
			}
		}
	}
	return frequencies
}

func populateAntinodesPlan(antennasPlan, antinodesPlan [][]rune, all bool) {
	frequencies := findAllFrequencies(antennasPlan)
	for _, positions := range frequencies {
		for _, fromPosition := range positions {
			for _, toPosition := range positions {
				delta := Position{toPosition.X - fromPosition.X, toPosition.Y - fromPosition.Y}
				populateAntinodesPlanFromPosition(antinodesPlan, toPosition, delta, all)
			}
		}
	}
}

func populateAntinodesPlanFromPosition(antinodesPlan [][]rune, position, delta Position, all bool) {
	for {
		if delta.X == 0 && delta.Y == 0 {
			if all {
				antinodesPlan[position.Y][position.X] = '#'
			}
			return
		}
		position.X += delta.X
		position.Y += delta.Y
		if position.X < 0 || position.Y < 0 || position.X >= len(antinodesPlan[0]) || position.Y >= len(antinodesPlan) {
			return
		}
		antinodesPlan[position.Y][position.X] = '#'
		if !all {
			return
		}
	}
}

func makeAntinodesPlan(antennasPlan [][]rune) [][]rune {
	antinodesPlan := make([][]rune, len(antennasPlan))
	for y, _ := range antennasPlan {
		antinodesPlan[y] = make([]rune, len(antennasPlan[y]))
		for x, _ := range antennasPlan[y] {
			antinodesPlan[y][x] = '.'
		}
	}
	return antinodesPlan
}

func countUniqueAntinodes(antinodesPlan [][]rune) int {
	count := 0
	for y, _ := range antinodesPlan {
		for x, _ := range antinodesPlan[y] {
			if antinodesPlan[y][x] == '#' {
				count++
			}
		}
	}
	return count
}
