package aoc2024

import (
	"fmt"

	_ "github.com/michelprogram/adventofcode/aoc_2024/day01"
	_ "github.com/michelprogram/adventofcode/aoc_2024/day02"
	_ "github.com/michelprogram/adventofcode/aoc_2024/day03"
	_ "github.com/michelprogram/adventofcode/aoc_2024/day04"
	_ "github.com/michelprogram/adventofcode/aoc_2024/day05"
	_ "github.com/michelprogram/adventofcode/aoc_2024/day06"
	_ "github.com/michelprogram/adventofcode/aoc_2024/day07"
	_ "github.com/michelprogram/adventofcode/aoc_2024/day08"
	_ "github.com/michelprogram/adventofcode/aoc_2024/day09"
	_ "github.com/michelprogram/adventofcode/aoc_2024/day10"
	_ "github.com/michelprogram/adventofcode/aoc_2024/day11"
	_ "github.com/michelprogram/adventofcode/aoc_2024/day12"
	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
)

type Aoc struct {
}

var _ utils.Code = (*Aoc)(nil)

func (a Aoc) Execute(data []byte, part, day int) (any, error) {

	if part == 1 {
		if challenge, ok := registry.GetChallenge(day); ok {
			return challenge.Part1(data)
		}
		return nil, fmt.Errorf("day %d doesn't register", day)
	}

	if challenge, ok := registry.GetChallenge(day); ok {
		return challenge.Part2(data)
	}
	return nil, fmt.Errorf("day %d doesn't register", day)

}
