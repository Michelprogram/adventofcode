package day12_test

import (
	"testing"

	"github.com/michelprogram/adventofcode/aoc_2024/day12"
	"github.com/michelprogram/adventofcode/utils"
)

type DataProvider struct {
	Inputs   []byte
	Expected int
}

type SideProvider struct {
	Inputs   []byte
	Expected map[string]int
}

func TestRunner_Part1(t *testing.T) {

	var runner day12.Runner

	providers := []DataProvider{
		{
			Inputs:   []byte("AAAA\nBBCD\nBBCC\nEEEC"),
			Expected: 140,
		},
		{
			Inputs:   []byte("OOOOO\nOXOXO\nOOOOO\nOXOXO\nOOOOO"),
			Expected: 772,
		},
		{
			Inputs:   []byte("RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE"),
			Expected: 1930,
		},
	}

	for _, provider := range providers {
		res, _ := runner.Part1(provider.Inputs)

		if res != provider.Expected {
			t.Fatalf("Should return %d instead %d\n", provider.Expected, res)
		}
	}
}

func TestRunner_Sides(t *testing.T) {

	providers := []SideProvider{
		{
			Inputs: []byte("AAAA\nBBCD\nBBCC\nEEEC"),
			Expected: map[string]int{
				"A": 4,
				"B": 4,
				"C": 8,
				"D": 4,
				"E": 4,
			},
		},
		{
			Inputs: []byte("EEEEE\nEXXXX\nEEEEE\nEXXXX\nEEEEE"),
			Expected: map[string]int{
				"E": 12,
				"X": 4,
			},
		},
	}

	for _, provider := range providers {

		garden := day12.NewGarden(provider.Inputs)

		visited := make(map[utils.Point[rune]]struct{})

		for _, plant := range garden.Plants {

			res, _ := garden.FindRegionArea(plant, visited)

			if res != nil {
				expected := garden.ComputeSides(*res)
				if expected != provider.Expected[string(plant.Value)] {
					t.Fatalf("Should return %d for letter %s instead %d\n", provider.Expected[string(plant.Value)], string(plant.Value), expected)
				}

			}

		}
	}
}

func TestRunner_Part2(t *testing.T) {

	var runner day12.Runner

	providers := []DataProvider{
		{
			Inputs:   []byte("AAAA\nBBCD\nBBCC\nEEEC"),
			Expected: 80,
		},
		{
			Inputs:   []byte("OOOOO\nOXOXO\nOOOOO\nOXOXO\nOOOOO"),
			Expected: 436,
		},
		{
			Inputs:   []byte("AAAAAA\nAAABBA\nAAABBA\nABBAAA\nABBAAA\nAAAAAA"),
			Expected: 368,
		},
		{
			Inputs:   []byte("RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE"),
			Expected: 1206,
		},
	}

	for _, provider := range providers {
		res, _ := runner.Part2(provider.Inputs)

		if res != provider.Expected {
			t.Fatalf("Should return %d instead %d\n", provider.Expected, res)
		}
	}
}
