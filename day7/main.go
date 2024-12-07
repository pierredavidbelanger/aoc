package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	day := "day7"
	part1ex := part1(parseInput(day + "/input-ex.txt"))
	part1exExpected := 3749
	log.Printf("part1 (example): %d\n", part1ex)
	if part1ex != part1exExpected {
		log.Fatalf("expecting %d\n", part1exExpected)
	}

	part1puzzle := part1(parseInput(day + "/input.txt"))
	log.Printf("part1 (puzzle): %d\n", part1puzzle)

	part2ex := part2(parseInput(day + "/input-ex.txt"))
	part2exExpected := 11387
	log.Printf("part2 (example): %d\n", part2ex)
	if part2ex != part2exExpected {
		log.Fatalf("expecting %d\n", part2exExpected)
	}

	part2puzzle := part2(parseInput(day + "/input.txt"))
	log.Printf("part2 (puzzle): %d\n", part2puzzle)
}

func parseInput(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileBuf := bufio.NewReader(file)
	equations := make([][]int, 0)
	for {
		line, err := fileBuf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		parts := strings.Split(strings.TrimSpace(line), ":")
		p0, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			log.Fatal(err)
		}
		equation := make([]int, 0)
		equation = append(equation, p0)
		parts = strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, part := range parts {
			p, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				log.Fatal(err)
			}
			equation = append(equation, p)
		}
		equations = append(equations, equation)
	}
	return equations
}

func part1(equations [][]int) int {
	sum := 0
	for _, equation := range equations {
		if testEquation([]rune{'+', '*'}, equation) {
			sum += equation[0]
		}
	}
	return sum
}

func part2(equations [][]int) int {
	sum := 0
	for _, equation := range equations {
		if testEquation([]rune{'+', '*', '|'}, equation) {
			sum += equation[0]
		}
	}
	return sum
}

func testEquation(possibleOps []rune, equation []int) bool {
	if testEquationRecursive(possibleOps, equation[0], equation[1:], []rune{}) {
		return true
	}
	return false
}

func testEquationRecursive(possibleOps []rune, expected int, numbers []int, ops []rune) bool {
	if len(ops) == len(numbers)-1 {
		total := numbers[0]
		for i := 0; i < len(numbers)-1; i++ {
			if ops[i] == '+' {
				total += numbers[i+1]
			} else if ops[i] == '*' {
				total *= numbers[i+1]
			} else if ops[i] == '|' {
				n, err := strconv.Atoi(fmt.Sprintf("%d%d", total, numbers[i+1]))
				if err != nil {
					log.Fatal(err)
				}
				total = n
			}
		}
		return total == expected
	}
	for _, possibleOp := range possibleOps {
		nops := make([]rune, 0)
		nops = append(nops, ops...)
		nops = append(nops, possibleOp)
		if testEquationRecursive(possibleOps, expected, numbers, nops) {
			return true
		}
	}
	return false
}
