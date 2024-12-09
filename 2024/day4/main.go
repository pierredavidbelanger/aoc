package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	wordSearch := parseInput()
	part1(wordSearch)
	part2(wordSearch)
}

func parseInput() [][]rune {

	wordSearch := make([][]rune, 0)

	inputFile, err := os.Open("2024/day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer inputFile.Close()
	inputBuffer := bufio.NewReader(inputFile)

	for {
		s, err := inputBuffer.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		wordSearch = append(wordSearch, []rune(strings.TrimSpace(s)))
	}

	return wordSearch
}

func found(wordSearch [][]rune, w, h, x, y, dx, dy int, word []rune, l int) int {
	//log.Printf("w=%d, h=%d, x=%d, y=%d, dx=%d, dy=%d, word=%s, l=%d", w, h, x, y, dx, dy, string(word), l)
	if l == 0 {
		return 1
	}
	if x < 0 || x >= w || y < 0 || y >= h {
		return 0
	}
	if wordSearch[x][y] != word[0] {
		return 0
	}
	return found(wordSearch, w, h, x+dx, y+dy, dx, dy, word[1:], l-1)
}

func foundAt(wordSearch [][]rune, w, h, x, y int, word []rune, l int) int {
	nb := 0
	nb += found(wordSearch, w, h, x, y, -1, 0, word, l)
	nb += found(wordSearch, w, h, x, y, 1, 0, word, l)
	nb += found(wordSearch, w, h, x, y, 0, 1, word, l)
	nb += found(wordSearch, w, h, x, y, 0, -1, word, l)
	nb += found(wordSearch, w, h, x, y, -1, -1, word, l)
	nb += found(wordSearch, w, h, x, y, -1, 1, word, l)
	nb += found(wordSearch, w, h, x, y, 1, -1, word, l)
	nb += found(wordSearch, w, h, x, y, 1, 1, word, l)
	return nb
}

func part1(wordSearch [][]rune) {

	w := len(wordSearch[0])
	h := len(wordSearch)

	nb := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			nb += foundAt(wordSearch, w, h, x, y, []rune("XMAS"), 4)
		}
	}

	log.Printf("found: %d\n", nb)
}

func part2(wordSearch [][]rune) {

	w := len(wordSearch[0])
	h := len(wordSearch)

	nb := 0

	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			if wordSearch[x][y] == 'A' {
				if wordSearch[x-1][y-1] == 'M' && wordSearch[x+1][y+1] == 'S' || wordSearch[x-1][y-1] == 'S' && wordSearch[x+1][y+1] == 'M' {
					if wordSearch[x+1][y-1] == 'M' && wordSearch[x-1][y+1] == 'S' || wordSearch[x+1][y-1] == 'S' && wordSearch[x-1][y+1] == 'M' {
						nb++
					}
				}
			}
		}
	}

	log.Printf("found: %d\n", nb)
}
