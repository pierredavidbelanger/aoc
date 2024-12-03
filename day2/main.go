package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := parseReports()
	part1(reports)
	part2(reports)
}

func parseReports() [][]int {

	inputFile, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer inputFile.Close()
	inputBuffer := bufio.NewReader(inputFile)

	reports := make([][]int, 0)

	for {

		line, err := inputBuffer.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		report := make([]int, 0)

		levels := strings.Split(line, " ")
		for _, levelString := range levels {
			level, err := strconv.Atoi(strings.TrimSpace(levelString))
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, level)
		}

		reports = append(reports, report)
	}

	return reports
}

func isSafe(report []int) bool {
	dir := 0
	for i := 1; i < len(report); i++ {
		d := report[i] - report[i-1]
		if d == 0 || d < -3 || d > 3 {
			return false
		}
		if i == 1 {
			dir = d
		} else if (dir < 0 && d > 0) || (dir > 0 && d < 0) {
			return false
		}
	}
	return true
}

func part1(reports [][]int) {

	safe := 0

	for _, report := range reports {
		if isSafe(report) {
			safe++
		}
	}

	log.Printf("safe: %d\n", safe)
}

func part2(reports [][]int) {

	safe := 0

	for _, report := range reports {
		if isSafe(report) {
			safe++
		} else {
			for i := 0; i < len(report); i++ {
				reportTest := make([]int, 0)
				reportTest = append(reportTest, report[:i]...)
				reportTest = append(reportTest, report[i+1:]...)
				if isSafe(reportTest) {
					safe++
					break
				}
			}
		}
	}

	log.Printf("safe: %d\n", safe)
}
