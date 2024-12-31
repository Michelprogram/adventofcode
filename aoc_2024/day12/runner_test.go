package day12_test

import (
	"github.com/michelprogram/adventofcode/aoc_2024/day12"
	"testing"
)

type DataProvider struct {
	Inputs   []byte
	Expected int
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
