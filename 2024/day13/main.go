package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type XY struct {
	X, Y int64
}

type Machine struct {
	ButtonA, ButtonB, Prize XY
	Token                   float64
}

func main() {

	day := "2024/day13"

	part1ex := part1(readInput(day + "/input-ex.txt"))
	part1exExpected := int64(480)
	log.Printf("part1 (example): %d\n", part1ex)
	if part1ex != part1exExpected {
		log.Fatalf("expecting %d\n", part1exExpected)
	}

	part1puzzle := part1(readInput(day + "/input.txt"))
	log.Printf("part1 (puzzle): %d\n", part1puzzle)

	part2ex := part2(readInput(day + "/input-ex.txt"))
	part2exExpected := int64(875318608908)
	log.Printf("part2 (example): %d\n", part2ex)
	if part2ex != part2exExpected {
		log.Fatalf("expecting %d\n", part2exExpected)
	}

	part2puzzle := part2(readInput(day + "/input.txt"))
	log.Printf("part2 (puzzle): %d\n", part2puzzle)
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

func part1(input []string) int64 {
	sum := int64(0)
	machines := inputToMachines(input)
	for i, _ := range machines {
		token := genSolveParse(&machines[i], "a <= 100;\nb <= 100;\n")
		if token > -1 {
			sum += token
		}
	}
	return sum
}

func part2(input []string) int64 {
	sum := int64(0)
	machines := inputToMachines(input)
	for i, _ := range machines {
		machines[i].Prize.X += 10000000000000
		machines[i].Prize.Y += 10000000000000
		token := genSolveParse(&machines[i], "")
		if token > -1 {
			sum += token
		}
	}
	return sum
}

func inputToMachines(input []string) []Machine {
	xyPattern := regexp.MustCompile("X[+=](\\d+), Y[+=](\\d+)")
	machines := make([]Machine, 0, len(input)/4)
	for i := 0; i < len(input); i += 4 {
		machine := Machine{}
		ba := xyPattern.FindAllStringSubmatch(input[i], -1)
		machine.ButtonA.X, _ = strconv.ParseInt(ba[0][1], 10, 64)
		machine.ButtonA.Y, _ = strconv.ParseInt(ba[0][2], 10, 64)
		bb := xyPattern.FindAllStringSubmatch(input[i+1], -1)
		machine.ButtonB.X, _ = strconv.ParseInt(bb[0][1], 10, 64)
		machine.ButtonB.Y, _ = strconv.ParseInt(bb[0][2], 10, 64)
		p := xyPattern.FindAllStringSubmatch(input[i+2], -1)
		machine.Prize.X, _ = strconv.ParseInt(p[0][1], 10, 64)
		machine.Prize.Y, _ = strconv.ParseInt(p[0][2], 10, 64)
		machines = append(machines, machine)
	}
	return machines
}

func genSolveParse(machine *Machine, constraints string) int64 {
	problem := genProblem(*machine, constraints)
	solution, err := solveProblem(problem)
	if err != nil {
		return -1
	}
	parseSolution(machine, solution)
	tokenInt, tokenFrac := math.Modf(machine.Token)
	if tokenFrac != 0 {
		return -1
	}
	return int64(tokenInt)
}

func genProblem(machine Machine, constraints string) string {
	return fmt.Sprintf(
		"min: 3a + b;\n"+constraints+"%da + %db = %d;\n%da + %db = %d;\nint a, b;\n",
		machine.ButtonA.X, machine.ButtonB.X, machine.Prize.X,
		machine.ButtonA.Y, machine.ButtonB.Y, machine.Prize.Y,
	)
}

func solveProblem(problem string) (string, error) {
	cmd := exec.Command("2024/day13/lp_solve.exe", "-S1")
	cmd.Stdin = strings.NewReader(problem)
	out := bytes.Buffer{}
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func parseSolution(machine *Machine, solution string) {
	objPattern := regexp.MustCompile("Value of objective function: ([\\w.]+)")
	obj := objPattern.FindAllStringSubmatch(solution, -1)
	machine.Token, _ = strconv.ParseFloat(obj[0][1], 64)
}
