package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	day := "2023/day2"

	part1ex := part1(readInput(day + "/input-ex.txt"))
	part1exExpected := 8
	log.Printf("part1 (example): %d\n", part1ex)
	if part1ex != part1exExpected {
		log.Fatalf("expecting %d\n", part1exExpected)
	}

	part1puzzle := part1(readInput(day + "/input.txt"))
	log.Printf("part1 (puzzle): %d\n", part1puzzle)

	part2ex := part2(readInput(day + "/input-ex.txt"))
	part2exExpected := 2286
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
	return strings.Split(string(data), "\n")
}

func part1(input []string) int {
	bag := map[string]int{"red": 12, "green": 13, "blue": 14}
	records := compact(parseInput(input))
	sum := 0
	for id, sets := range records {
		if isGamePossible(sets, bag) {
			sum += id
		}
	}
	return sum
}

func part2(input []string) int {
	records := compact(parseInput(input))
	sum := 0
	for _, sets := range records {
		sum += getPower(sets)
	}
	return sum
}

func parseInput(input []string) map[int][]map[string]int {
	records := make(map[int][]map[string]int, len(input))
	for _, record := range input {
		gameAndSets := strings.Split(record, ":")
		game := strings.Split(gameAndSets[0], " ")
		gameId, err := strconv.Atoi(strings.TrimSpace(game[1]))
		if err != nil {
			log.Fatal(err)
		}
		sets := strings.Split(gameAndSets[1], ";")
		records[gameId] = make([]map[string]int, len(sets))
		for s, set := range sets {
			cubes := strings.Split(strings.TrimSpace(set), ",")
			records[gameId][s] = make(map[string]int, len(cubes))
			for _, cube := range cubes {
				nbAndColor := strings.Split(strings.TrimSpace(cube), " ")
				nb, err := strconv.Atoi(strings.TrimSpace(nbAndColor[0]))
				if err != nil {
					log.Fatal(err)
				}
				color := strings.TrimSpace(nbAndColor[1])
				records[gameId][s][color] += nb
			}
		}
	}
	return records
}

func compact(records map[int][]map[string]int) map[int]map[string]int {
	crecords := make(map[int]map[string]int, len(records))
	for id, sets := range records {
		crecords[id] = make(map[string]int, 5)
		for _, set := range sets {
			for color, nb := range set {
				if nb > crecords[id][color] {
					crecords[id][color] = nb
				}
			}
		}
	}
	return crecords
}

func isGamePossible(sets map[string]int, bag map[string]int) bool {
	for color, nb := range sets {
		if nb > bag[color] {
			return false
		}
	}
	return true
}

func getPower(sets map[string]int) int {
	power := -1
	for _, nb := range sets {
		if power == -1 {
			power = nb
		} else {
			power *= nb
		}
	}
	return power
}
