package main

import "testing"

func TestPerf(t *testing.T) {
	input := parseInput("input.txt")
	stones := inputToStones(input)
	blink(stones, 50)
}
