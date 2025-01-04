package day13_test

import (
	"testing"

	"github.com/michelprogram/adventofcode/aoc_2024/day13"
)

type DataProvider struct {
	Inputs   []byte
	Expected int
}

func Test1(t *testing.T) {
	var runner day13.Runner

	providers := []DataProvider{
		{
			Inputs:   []byte("Button A: X+94, Y+34\nButton B: X+22, Y+67\nPrize: X=8400, Y=5400\n\n"),
			Expected: 280,
		},
	}

	for _, provider := range providers {
		res, _ := runner.Part1(provider.Inputs)

		if res != provider.Expected {
			t.Fatalf("Should return %d instead %d\n", provider.Expected, res)

		}

	}

}
