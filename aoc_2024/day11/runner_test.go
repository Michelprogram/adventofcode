package day11_test

import (
	"testing"

	"github.com/michelprogram/adventofcode/aoc_2024/day11"
)

type DataProvider struct {
	Inputs    []int
	Iteration int
	Expected  int
}

func TestWorker(t *testing.T) {
	var runner day11.Runner

	providers := []DataProvider{
		{
			Inputs:    []int{125},
			Iteration: 78,
			Expected:  22,
		},
		{
			Inputs:    []int{125, 17},
			Iteration: 6,
			Expected:  22,
		},
		{
			Inputs:    []int{125, 17},
			Iteration: 25,
			Expected:  55312,
		},
		{
			Inputs:    []int{0, 1, 10, 99, 999},
			Iteration: 1,
			Expected:  7,
		},
	}

	for _, provider := range providers {
		res := runner.Worker(provider.Inputs, provider.Iteration)

		if provider.Expected != res {
			t.Fatalf("Should return %d instead %d\n", provider.Expected, res)
		}
	}
}

/*
func TestPart1(t *testing.T) {
	var runner day11.Runner

	providers := []DataProvider{

		{
			Inputs:   []byte("125 17"),
			Expected: 22,
		},
		{
			Inputs:   []byte("0 1 10 99 999"),
			Expected: 7,
		},
	}

	for _, provider := range providers {
		res, err := runner.Part1(provider.Inputs)

		if err != nil {
			t.Fatalf("There shouldn't be error: %s\n", err.Error())
		}

		if res != provider.Expected {
			t.Fatalf("Should return %d instead %d\n", provider.Expected, res)
		}
	}
}
*/
