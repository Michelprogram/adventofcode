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

func TestRunner_WorkerWithMemoization(t *testing.T) {
	var runner day11.Runner

	providers := []DataProvider{

		{
			Inputs:    []int{125, 17},
			Iteration: 6,
			Expected:  22,
		},
		{
			Inputs:    []int{125, 17},
			Iteration: 75,
			Expected:  22,
		},
	}

	for _, provider := range providers {
		res := runner.WorkerWithMemoization(provider.Inputs, provider.Iteration)

		if res != provider.Expected {
			t.Fatalf("Should return %d instead %d\n", provider.Expected, res)
		}
	}
}
