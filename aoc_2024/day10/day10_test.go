package day10_test

import (
	"testing"

	"github.com/michelprogram/adventofcode/aoc_2024/day10"
)

type DataProvider struct {
	Inputs   []byte
	Expected int
}

func TestPath(t *testing.T) {

	inputs := []byte("0123\n9234\n8765\n9876")

	graph, err := day10.NewGraph(inputs)

	if err != nil {
		t.Fatalf("There shouldn't be error: %s\n", err.Error())
	}

	if len(graph.Started) != 1 {
		t.Fatalf("Should return 1 instead %d\n", len(graph.Started))
	}

	paths := graph.FindPath(*graph.Started[0])

	if paths != 2 {
		t.Fatalf("Should return 2 instead %d\n", paths)
	}
}

func TestPart1(t *testing.T) {

	var runner day10.Runner

	providers := []DataProvider{
		{
			Inputs:   []byte("1190119\n1111198\n1112117\n6543456\n7651987\n8761111\n9871111"),
			Expected: 4,
		},
		{
			Inputs:   []byte("9990999\n9991999\n9992999\n6543456\n7111117\n8111118\n9111199"),
			Expected: 2,
		},
		{
			Inputs:   []byte("89010123\n78121874\n87430965\n96549874\n45678903\n32019012\n01329801\n10456732"),
			Expected: 36,
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

func TestPart2(t *testing.T) {

	var runner day10.Runner

	providers := []DataProvider{
		{
			Inputs:   []byte("9999909\n9943219\n9959929\n9965439\n9971149\n1187651\n1191119"),
			Expected: 3,
		},
		{
			Inputs:   []byte("1190339\n3331398\n9992997\n6543456\n7651987\n8761111\n9871111"),
			Expected: 13,
		},
		{
			Inputs:   []byte("012345\n123456\n234567\n345678\n416789\n567891"),
			Expected: 227,
		},
	}

	for _, provider := range providers {

		res, err := runner.Part2(provider.Inputs)

		if err != nil {
			t.Fatalf("There shouldn't be error: %s\n", err.Error())
		}

		if res != provider.Expected {
			t.Fatalf("Should return %d instead %d\n", provider.Expected, res)
		}
	}

}
