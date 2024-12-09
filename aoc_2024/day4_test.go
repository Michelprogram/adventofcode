package aoc2024

import (
	"testing"
)

func TestDay4Horizontal(t *testing.T) {

	inputs := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"MXAMFMSXSA",
	}

	day := NewDay4(inputs)

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

	day := NewDay4(inputs)

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

	day := NewDay4(inputs)

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

	day := NewDay4(inputs)

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

	day := NewDay4(inputs)

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

	day := NewDay4(inputs)

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
