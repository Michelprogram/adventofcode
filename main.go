package main

import (
	"flag"

	aoc2024 "github.com/michelprogram/adventofcode/aoc_2024"
	"github.com/michelprogram/adventofcode/utils"
)

func main() {

	var day int
	var year int
	var part int
	var generator bool

	flag.IntVar(&year, "year", 2024, "select year")
	flag.IntVar(&day, "day", 12, "select day")
	flag.IntVar(&part, "part", 2, "part could be either 1 or 2")
	flag.BoolVar(&generator, "generator", false, "generate files structure")

	flag.Parse()

	if generator {
		err := utils.GenerateFiles(day, year)
		if err != nil {
			panic(err)
		}

		return

	}

	aocs := map[int]utils.Code{
		2024: aoc2024.Aoc{},
	}

	err := utils.RunAoc(part, day, year, aocs)
	if err != nil {
		panic(err)
	}

}
