package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Vec struct {
	X, Y int
}

type Robot struct {
	Pos Vec
	Vel Vec
}

func main() {

	println(-2 % 10)

	day := "2024/day14"

	part1ex := part1(readInput(day+"/input-ex.txt"), Vec{11, 7}, 100)
	part1exExpected := 12
	log.Printf("part1 (example): %d\n", part1ex)
	if part1ex != part1exExpected {
		log.Fatalf("expecting %d\n", part1exExpected)
	}

	part1puzzle := part1(readInput(day+"/input.txt"), Vec{101, 103}, 100)
	log.Printf("part1 (puzzle): %d\n", part1puzzle)

	part2ex := part2(readInput(day + "/input-ex.txt"))
	part2exExpected := 0
	log.Printf("part2 (example): %d\n", part2ex)
	if part2ex != part2exExpected {
		log.Fatalf("expecting %d\n", part2exExpected)
	}

	part2puzzle := part2(readInput(day + "/input.txt"))
	log.Printf("part2 (puzzle): %d\n", part2puzzle)
}

func part1(input []string, dim Vec, time int) int {
	robots := inputToRobots(input)
	moveRobots(robots, dim, time)
	q1, q2, q3, q4 := robotCountPerQuadrant(robots, dim)
	return q1 * q2 * q3 * q4
}

func part2(input []string) int {
	return 0
}

func readInput(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	input := make([]string, len(lines))
	for i, line := range lines {
		input[i] = strings.TrimSpace(line)
	}
	return input
}

func inputToRobots(input []string) []Robot {
	robotPattern := regexp.MustCompile("p=(\\d+),(\\d+) v=(-?\\d+),(-?\\d+)")
	robots := make([]Robot, 0, len(input))
	for _, line := range input {
		robot := Robot{}
		matches := robotPattern.FindStringSubmatch(line)
		robot.Pos.X = parseInt(matches[1])
		robot.Pos.Y = parseInt(matches[2])
		robot.Vel.X = parseInt(matches[3])
		robot.Vel.Y = parseInt(matches[4])
		robots = append(robots, robot)
	}
	return robots
}

func parseInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 32)
	return int(i)
}

func moveRobots(robots []Robot, dim Vec, time int) {
	for r, _ := range robots {
		robots[r].Pos.X += robots[r].Vel.X * time
		robots[r].Pos.X = robots[r].Pos.X % dim.X
		if robots[r].Pos.X < 0 {
			robots[r].Pos.X += dim.X
		}
		robots[r].Pos.Y += robots[r].Vel.Y * time
		robots[r].Pos.Y = robots[r].Pos.Y % dim.Y
		if robots[r].Pos.Y < 0 {
			robots[r].Pos.Y += dim.Y
		}
	}
}

func robotCountPerQuadrant(robots []Robot, dim Vec) (int, int, int, int) {
	q1, q2, q3, q4 := 0, 0, 0, 0
	hx, hy := dim.X/2, dim.Y/2
	for _, robot := range robots {
		if robot.Pos.X < hx && robot.Pos.Y < hy {
			q1++
		} else if robot.Pos.X > hx && robot.Pos.Y < hy {
			q2++
		} else if robot.Pos.X < hx && robot.Pos.Y > hy {
			q3++
		} else if robot.Pos.X > hx && robot.Pos.Y > hy {
			q4++
		}
	}
	return q1, q2, q3, q4
}
