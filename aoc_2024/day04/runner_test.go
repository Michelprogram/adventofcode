package day4_test

import (
	"testing"

	"github.com/michelprogram/adventofcode/aoc_2024/day04"
)

func TestDay4Horizontal(t *testing.T) {

	inputs := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"MXAMFMSXSA",
	}

	day := day4.NewDay4(inputs)

	for y, line := range inputs {
		for x, letter := range line {
			if letter == rune('X') {
				day.FindHorizontal(x, y)
			}
		}
	}

	if day.Counter != 2 {
		t.Fatalf("Sould return 2 instead %d\n", day.Counter)
	}
}

func TestDay4Vertical(t *testing.T) {

	inputs := []string{
		"MXMSFXMASS",
		"MMAMXXSMSA",
		"MAAMFMXFSM",
		"MSAMFMSFSX",
	}

	day := day4.NewDay4(inputs)

	for y, line := range inputs {
		for x, letter := range line {
			if letter == rune('X') {
				day.FindVertical(x, y)
			}
		}
	}
	if day.Counter != 2 {
		t.Fatalf("Sould return 2 instead %d\n", day.Counter)
	}
}

func TestDay4CrossTopLeft(t *testing.T) {

	inputs := []string{
		"MMAMMSMMXM",
		"MMXMMMAMMM",
		"MMMMXMMMMM",
		"MXMMMMMMXM",
	}

	day := day4.NewDay4(inputs)

	for y, line := range inputs {
		for x, letter := range line {
			if letter == rune('X') {
				day.FindCross(x, y)
			}
		}
	}

	if day.Counter != 1 {
		t.Fatalf("Sould return 1 instead %d\n", day.Counter)
	}
}

func TestDay4CrossTopRight(t *testing.T) {

	inputs := []string{
		"MMAMSMMMXA",
		"MMXAMMAMMM",
		"MMMMXMMXMM",
		"MXMMXMMMXM",
	}

	day := day4.NewDay4(inputs)

	for y, line := range inputs {
		for x, letter := range line {
			if letter == rune('X') {
				day.FindCross(x, y)
			}
		}
	}

	if day.Counter != 1 {
		t.Fatalf("Sould return 1 instead %d\n", day.Counter)
	}
}

func TestDay4CrossBotLeft(t *testing.T) {

	inputs := []string{
		"MMXMXMXMMM",
		"MMXMMMMMXM",
		"MMAMMMMMAM",
		"ASMMMMMSMM",
	}

	day := day4.NewDay4(inputs)

	for y, line := range inputs {
		for x, letter := range line {
			if letter == rune('X') {
				day.FindCross(x, y)
			}
		}
	}

	if day.Counter != 1 {
		t.Fatalf("Sould return 1 instead %d\n", day.Counter)
	}
}

func TestDay4CrossBotRight(t *testing.T) {

	inputs := []string{
		"MMXMXMXMMM",
		"MMMMMMMMXM",
		"MXMMMMAMAM",
		"MXMMMMMSMM",
	}

	day := day4.NewDay4(inputs)

	for y, line := range inputs {
		for x, letter := range line {
			if letter == rune('X') {
				day.FindCross(x, y)
			}
		}
	}

	if day.Counter != 1 {
		t.Fatalf("Sould return 1 instead %d\n", day.Counter)
	}
}
