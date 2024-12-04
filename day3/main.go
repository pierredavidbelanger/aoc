package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part1(getInput1())
	log.Println()
	part2(getInput2())
}

func getInput1() string {
	data, err := os.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func getInput2() string {
	data, err := os.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func part1(input string) {
	re, err := regexp.Compile("mul\\((\\d{1,3}?),(\\d{1,3}?)\\)")
	if err != nil {
		log.Fatal(err)
	}
	matchs := re.FindAllStringSubmatch(input, -1)
	total := 0
	for _, match := range matchs {
		n1, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		n2, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s = %d\n", match[0], n1*n2)
		total += n1 * n2
	}
	log.Printf("total %d\n", total)
}

func part2(input string) {
	re, err := regexp.Compile("(mul\\((\\d{1,3}?),(\\d{1,3}?)\\))|(do\\(\\))|(don't\\(\\))")
	if err != nil {
		log.Fatal(err)
	}
	matchs := re.FindAllStringSubmatch(input, -1)
	total := 0
	do := true
	for _, match := range matchs {
		if match[0] == "do()" {
			do = true
			log.Printf("%s\n", match[0])
		} else if match[0] == "don't()" {
			do = false
			log.Printf("%s\n", match[0])
		} else {
			n1, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatal(err)
			}
			n2, err := strconv.Atoi(match[3])
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%s = %d (do=%v)\n", match[1], n1*n2, do)
			if do {
				total += n1 * n2
			}
		}
	}
	log.Printf("total %d\n", total)
}
