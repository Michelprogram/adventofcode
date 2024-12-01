package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	aoc2024 "github.com/michelprogram/adventofcode/aoc_2024"
	"github.com/michelprogram/adventofcode/utils"
)

func main() {

	var day int
	var year int
	var part int

	flag.IntVar(&year, "year", 2023, "select year")
	flag.IntVar(&day, "day", 1, "select day")
	flag.IntVar(&part, "part", 1, "part could be either 1 or 2")

	flag.Parse()

	data, err := utils.FecthDataSet(year, day)

	if err != nil {
		log.Fatalf("Can't fetch data for day %d\n", day)
	}

	years := map[int]utils.Code{
		2024: aoc2024.Aoc{},
	}

	aoc, ok := years[year]

	if !ok {
		log.Fatalf("year %d doesn't exist\n", year)
	}

	start := time.Now()

	res, err := aoc.Execute(data, part, day)

	if err != nil {
		log.Fatalf("Error during advent code %d for day %d : %s\n", year, day, err)
	}

	fmt.Printf("Time : %v\n", time.Since(start))

	fmt.Printf("Resultat for day %d : %v\n", day, res)

}
