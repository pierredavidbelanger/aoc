package main

import (
	"log"
	"os"
	"strconv"
)

func main() {

	day := "2024/day9"

	part1ex := part1(parseInput(day + "/input-ex.txt"))
	part1exExpected := 1928
	log.Printf("part1 (example): %d\n", part1ex)
	if part1ex != part1exExpected {
		log.Fatalf("expecting %d\n", part1exExpected)
	}

	part1puzzle := part1(parseInput(day + "/input.txt"))
	log.Printf("part1 (puzzle): %d\n", part1puzzle)

	part2ex := part2(parseInput(day + "/input-ex.txt"))
	part2exExpected := 2858
	log.Printf("part2 (example): %d\n", part2ex)
	if part2ex != part2exExpected {
		log.Fatalf("expecting %d\n", part2exExpected)
	}

	part2puzzle := part2(parseInput(day + "/input.txt"))
	log.Printf("part2 (puzzle): %d\n", part2puzzle)
}

func parseInput(fileName string) string {
	input, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}

func part1(input string) int {
	df := inputToDenseFormat(input)
	sf := denseToSparseFormat(df)
	dsf := defragment1(sf)
	return checksum(dsf)
}

func part2(input string) int {
	df := inputToDenseFormat(input)
	sf := denseToSparseFormat(df)
	dsf := defragment2(sf)
	return checksum(dsf)
}

func inputToDenseFormat(input string) []int {
	df := make([]int, len(input))
	for i, c := range input {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal(err)
		}
		df[i] = n
	}
	return df
}

func denseToSparseFormat(df []int) []int {
	sf := make([]int, 0)
	id := 0
	for i, n := range df {
		nn := -1
		if i%2 == 0 {
			nn = id
			id++
		}
		for j := 0; j < n; j++ {
			sf = append(sf, nn)
		}
	}
	return sf
}

func defragment1(sf []int) []int {
	dsf := make([]int, len(sf))
	copy(dsf, sf)
	li := 0
	ri := len(dsf) - 1
	for li < ri {
		for li < ri && dsf[li] != -1 {
			li++
		}
		for ri > li && dsf[ri] == -1 {
			ri--
		}
		if li != ri {
			dsf[li], dsf[ri] = dsf[ri], dsf[li]
			li++
			ri--
		}
	}
	return dsf
}

func defragment2(sf []int) []int {
	dsf := make([]int, len(sf))
	copy(dsf, sf)
	ss, fb, fe := 0, len(dsf)-1, len(dsf)-1
	for ss < fb {
		for ss < fb && dsf[ss] != -1 {
			ss++
		}
		for fe > ss && dsf[fe] == -1 {
			fe--
		}
		fb = fe
		for fb > ss && dsf[fb-1] == dsf[fe] {
			fb--
		}
		fl := fe - fb + 1
		sb, se := ss, ss
		for se < fb {
			for sb < fb && dsf[sb] != -1 {
				sb++
			}
			se = sb
			for se < fb && dsf[se+1] == dsf[sb] {
				se++
			}
			sl := se - sb + 1
			if sl >= fl {
				for mi := 0; mi < fl; mi++ {
					dsf[sb+mi], dsf[fb+mi] = dsf[fb+mi], dsf[sb+mi]
				}
				break
			}
			sb, se = se+1, se+1
		}
		fb, fe = fb-1, fb-1
	}
	return dsf
}

func checksum(sf []int) int {
	sum := 0
	for i, n := range sf {
		if n != -1 {
			sum += i * n
		}
	}
	return sum
}
