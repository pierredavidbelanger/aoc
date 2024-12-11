package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func main() {

	day := "2023/day1"

	part1ex := part1(parseInput(day + "/input-ex.txt"))
	part1exExpected := 142
	log.Printf("part1 (example): %d\n", part1ex)
	if part1ex != part1exExpected {
		log.Fatalf("expecting %d\n", part1exExpected)
	}

	part1puzzle := part1(parseInput(day + "/input.txt"))
	log.Printf("part1 (puzzle): %d\n", part1puzzle)

	part2ex := part2(parseInput(day + "/input-ex2.txt"))
	part2exExpected := 281
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
		}
		input = append(input, strings.TrimSpace(line))
	}
	return input
}

func part1(calibrations []string) int {
	values := 0
	for _, calibration := range calibrations {
		values += getCalibrationValue(calibration, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"})
	}
	return values
}

func part2(calibrations []string) int {
	values := 0
	for _, calibration := range calibrations {
		values += getCalibrationValue(calibration, []string{"one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9"})
	}
	return values
}

func getCalibrationValue(calibration string, numbers []string) int {
	n1, _ := getLeftMostNumber(calibration, numbers)
	n2, _ := getRightMostNumber(calibration, numbers)
	return n1*10 + n2
}

func getLeftMostNumber(calibration string, numbers []string) (int, string) {
	idx, number := math.MaxInt, ""
	for _, n := range numbers {
		i := strings.Index(calibration, n)
		if i != -1 && i < idx {
			idx, number = i, n
		}
	}
	return asInt(number), number
}

func getRightMostNumber(calibration string, numbers []string) (int, string) {
	idx, number := -1, ""
	for _, n := range numbers {
		i := strings.LastIndex(calibration, n)
		if i != -1 && i > idx {
			idx, number = i, n
		}
	}
	return asInt(number), number
}

func asInt(number string) int {
	switch number {
	case "1", "one":
		return 1
	case "2", "two":
		return 2
	case "3", "three":
		return 3
	case "4", "four":
		return 4
	case "5", "five":
		return 5
	case "6", "six":
		return 6
	case "7", "seven":
		return 7
	case "8", "eight":
		return 8
	case "9", "nine":
		return 9
	}
	log.Fatalf("invalid number: %s", number)
	return 0
}
