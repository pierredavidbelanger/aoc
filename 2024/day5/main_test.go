package day5

import (
	"testing"
)

func TestPart1WithTest(t *testing.T) {
	rules, updates := parse("input-test.txt")
	correctUpdates, _ := findCorrectlyOrderedUpdates(rules, updates)
	sum := sumMiddlePages(correctUpdates)
	t.Logf("Sum: %d", sum)
	if sum != 143 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 143)
	}
}

func TestPart1WithData(t *testing.T) {
	rules, updates := parse("input.txt")
	correctUpdates, _ := findCorrectlyOrderedUpdates(rules, updates)
	sum := sumMiddlePages(correctUpdates)
	t.Logf("Sum: %d", sum)
}

func TestPart2WithTest(t *testing.T) {
	rules, updates := parse("input-test.txt")
	_, incorrectUpdates := findCorrectlyOrderedUpdates(rules, updates)
	fixUpdates(rules, &incorrectUpdates)
	sum := sumMiddlePages(incorrectUpdates)
	t.Logf("Sum: %d", sum)
	if sum != 123 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 143)
	}
}

func TestPart2WithData(t *testing.T) {
	rules, updates := parse("input.txt")
	_, incorrectUpdates := findCorrectlyOrderedUpdates(rules, updates)
	fixUpdates(rules, &incorrectUpdates)
	sum := sumMiddlePages(incorrectUpdates)
	t.Logf("Sum: %d", sum)
}
