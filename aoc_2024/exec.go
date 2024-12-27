package aoc2024

import (
	"fmt"

	_ "github.com/michelprogram/adventofcode/aoc_2024/day10"
	_ "github.com/michelprogram/adventofcode/aoc_2024/day11"
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
		return nil, fmt.Errorf("Day %d doesn't register", day)
	}

	if challenge, ok := registry.GetChallenge(day); ok {
		return challenge.Part2(data)
	}
	return nil, fmt.Errorf("Day %d doesn't register", day)

}
