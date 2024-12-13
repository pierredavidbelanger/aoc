package main

import (
	"log"
	"math"
	"os"
	"strings"
)

type XY struct {
	X, Y int
}

func main() {

	day := "2024/day12"

	part1ex := part1(readInput(day + "/input-ex.txt"))
	part1exExpected := 1930
	log.Printf("part1 (example): %d\n", part1ex)
	if part1ex != part1exExpected {
		log.Fatalf("expecting %d\n", part1exExpected)
	}

	part1puzzle := part1(readInput(day + "/input.txt"))
	log.Printf("part1 (puzzle): %d\n", part1puzzle)

	part2ex := part2(readInput(day + "/input-ex.txt"))
	part2exExpected := 1206
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
	garden := inputToGarden(input)
	regions := findRegions(garden)
	return regionsPriceForAreaAndPerimeter(garden, regions)
}

func part2(input []string) int {
	garden := inputToGarden(input)
	regions := findRegions(garden)
	return regionsPriceForAreaAndSides(garden, regions)
}

func inputToGarden(input []string) [][]rune {
	garden := make([][]rune, len(input))
	for i, line := range input {
		garden[i] = []rune(strings.TrimSpace(line))
	}
	return garden
}

func emptyGarden(g1 [][]rune) [][]rune {
	g2 := make([][]rune, len(g1))
	for y, row := range g1 {
		g2[y] = make([]rune, len(row))
		for x, _ := range row {
			g2[y][x] = -1
		}
	}
	return g2
}

func findRegions(garden [][]rune) [][]XY {
	regions := make([][]XY, 0)
	gardenFill := emptyGarden(garden)
	for y, row := range garden {
		for x, _ := range row {
			region := regionAt(garden, gardenFill, XY{x, y})
			if region != nil {
				regions = append(regions, region)
			}
		}
	}
	return regions
}

func regionAt(gardenOrig [][]rune, gardenFill [][]rune, pos XY) []XY {
	if gardenFill[pos.Y][pos.X] != -1 {
		return nil
	}
	return floodFill(gardenOrig, gardenFill, gardenOrig[pos.Y][pos.X], pos)
}

func floodFill(gardenOrig [][]rune, gardenFill [][]rune, id rune, start XY) []XY {
	region := make([]XY, 0)
	stack := make([]XY, 0)
	stack = append(stack, start)
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if gardenOrig[p.Y][p.X] == id && gardenFill[p.Y][p.X] == -1 {
			gardenFill[p.Y][p.X] = 1
			region = append(region, p)
			if p.Y+1 < len(gardenOrig) {
				stack = append(stack, XY{p.X, p.Y + 1})
			}
			if p.Y-1 >= 0 {
				stack = append(stack, XY{p.X, p.Y - 1})
			}
			if p.X-1 >= 0 {
				stack = append(stack, XY{p.X - 1, p.Y})
			}
			if p.X+1 < len(gardenOrig[0]) {
				stack = append(stack, XY{p.X + 1, p.Y})
			}
		}
	}
	return region
}

func regionsPriceForAreaAndPerimeter(garden [][]rune, regions [][]XY) int {
	price := 0
	for _, region := range regions {
		area, perimeter := regionAreaAndPerimeter(garden, region)
		price += area * perimeter
	}
	return price
}

func regionAreaAndPerimeter(garden [][]rune, region []XY) (int, int) {
	area, perimeter := len(region), 0
	id := plotAt(garden, region[0].X, region[0].Y)
	for _, xy := range region {
		if plotAt(garden, xy.X, xy.Y+1) != id {
			perimeter++
		}
		if plotAt(garden, xy.X, xy.Y-1) != id {
			perimeter++
		}
		if plotAt(garden, xy.X-1, xy.Y) != id {
			perimeter++
		}
		if plotAt(garden, xy.X+1, xy.Y) != id {
			perimeter++
		}
	}
	return area, perimeter
}

func regionsPriceForAreaAndSides(garden [][]rune, regions [][]XY) int {
	price := 0
	for _, region := range regions {
		area, sides := regionAreaAndSides(garden, region)
		price += area * sides
	}
	return price
}

func regionAreaAndSides(garden [][]rune, region []XY) (int, int) {
	area, sides := len(region), 0
	id := plotAt(garden, region[0].X, region[0].Y)
	minXY := XY{math.MaxInt, math.MaxInt}
	maxXY := XY{-1, -1}
	for _, xy := range region {
		minXY.X = min(xy.X, minXY.X)
		minXY.Y = min(xy.Y, minXY.Y)
		maxXY.X = max(xy.X, maxXY.X)
		maxXY.Y = max(xy.Y, maxXY.Y)
	}
	for y := minXY.Y; y <= maxXY.Y; y++ {
		side1, side2 := false, false
		for x := minXY.X; x <= maxXY.X; x++ {
			if inRegion(region, x, y) && plotAt(garden, x, y) == id {
				if plotAt(garden, x, y-1) != id {
					if !side1 {
						side1 = true
						sides++
					}
				} else {
					side1 = false
				}
				if plotAt(garden, x, y+1) != id {
					if !side2 {
						side2 = true
						sides++
					}
				} else {
					side2 = false
				}
			} else {
				side1 = false
				side2 = false
			}
		}
	}
	for x := minXY.X; x <= maxXY.X; x++ {
		side1, side2 := false, false
		for y := minXY.Y; y <= maxXY.Y; y++ {
			if inRegion(region, x, y) && plotAt(garden, x, y) == id {
				if plotAt(garden, x-1, y) != id {
					if !side1 {
						side1 = true
						sides++
					}
				} else {
					side1 = false
				}
				if plotAt(garden, x+1, y) != id {
					if !side2 {
						side2 = true
						sides++
					}
				} else {
					side2 = false
				}
			} else {
				side1 = false
				side2 = false
			}
		}
	}
	return area, sides
}

func plotAt(garden [][]rune, x, y int) rune {
	if x < 0 || y < 0 || x >= len(garden[0]) || y >= len(garden) {
		return -1
	}
	return garden[y][x]
}

func inRegion(region []XY, x, y int) bool {
	for _, xy := range region {
		if xy.X == x && xy.Y == y {
			return true
		}
	}
	return false
}
