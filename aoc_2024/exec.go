package aoc2024

import (
	"github.com/michelprogram/adventofcode/utils"
)

type Aoc struct {
}

var _ utils.Code = (*Aoc)(nil)

func (a Aoc) Execute(data []byte, part, day int) (any, error) {

	days := []utils.Challenge{
		Day1{},
		Day2{},
		Day3{},
		Day4{},
		Day5{},
		Day6{},
	}

	if part == 1 {
		return days[day-1].Part1(data)
	}

	return days[day-1].Part2(data)

}
