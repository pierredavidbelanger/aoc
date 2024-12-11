package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	day := "2024/day11"

	part1ex := part1(parseInput(day + "/input-ex.txt"))
	part1exExpected := 55312
	log.Printf("part1 (example): %d\n", part1ex)
	if part1ex != part1exExpected {
		log.Fatalf("expecting %d\n", part1exExpected)
	}

	part1puzzle := part1(parseInput(day + "/input.txt"))
	log.Printf("part1 (puzzle): %d\n", part1puzzle)

	part2puzzle := part2(parseInput(day + "/input.txt"))
	log.Printf("part2 (puzzle): %d\n", part2puzzle)
}

func parseInput(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileBuf := bufio.NewReader(file)
	line, err := fileBuf.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return line
}

func part1(input string) int {
	stones := inputToStones(input)
	return blink(stones, 25)
}

func part2(input string) int {
	stones := inputToStones(input)
	return blink(stones, 75)
}

func inputToStones(input string) []int {
	parts := strings.Split(input, " ")
	stones := make([]int, len(parts))
	for i, part := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			log.Fatal(err)
		}
		stones[i] = n
	}
	return stones
}

func blink(stones []int, times int) int {
	count := 0
	cache := make(map[string]int, 1000000)
	for _, stone := range stones {
		count += blinkRecursive(stone, times, cache)
	}
	return count
}

func blinkRecursive(stone int, times int, cache map[string]int) int {
	if times == 0 {
		return 1
	}
	k := fmt.Sprintf("%d-%d", stone, times)
	if count, ok := cache[k]; ok {
		return count
	}
	times--
	lh, rh := transform(stone)
	count := blinkRecursive(lh, times, cache)
	if rh > -1 {
		count += blinkRecursive(rh, times, cache)
	}
	cache[k] = count
	return count
}

func transform(stone int) (int, int) {
	if stone == 0 {
		return 1, -1
	}
	number := fmt.Sprintf("%d", stone)
	if len(number)%2 == 0 {
		lh, _ := strconv.Atoi(number[:len(number)/2])
		rh, _ := strconv.Atoi(number[len(number)/2:])
		return lh, rh
	}
	return stone * 2024, -1
}
