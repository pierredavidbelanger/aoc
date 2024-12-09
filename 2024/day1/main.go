package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func main() {

	l1, l2 := parseLists()

	part1(l1, l2)

	part2(l1, l2)
}

func parseLists() ([]int, []int) {

	inputFile, err := os.Open("2024/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer inputFile.Close()

	l1, l2 := make([]int, 0), make([]int, 0)
	for {
		i1, i2 := 0, 0
		_, err = fmt.Fscanf(inputFile, "%d   %d\n", &i1, &i2)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		l1 = append(l1, i1)
		l2 = append(l2, i2)
	}

	sort.Ints(l1)
	sort.Ints(l2)

	return l1, l2
}

func part1(l1, l2 []int) {

	td := 0
	for i := 0; i < len(l1); i++ {
		d := l1[i] - l2[i]
		if d < 0 {
			td -= d
		} else {
			td += d
		}
	}

	log.Printf("total distance: %d\n", td)
}

func part2(l1, l2 []int) {

	ss := 0

	for _, i1 := range l1 {
		c := 0
		for _, i2 := range l2 {
			if i2 > i1 {
				break
			}
			if i1 == i2 {
				c++
			}
		}
		ss += i1 * c
	}

	log.Printf("similarity score: %d\n", ss)
}
