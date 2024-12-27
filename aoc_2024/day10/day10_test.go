package day10_test

import (
	"testing"

	"github.com/michelprogram/adventofcode/aoc_2024/day10"
)

func Test1(t *testing.T) {

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

func Test2(t *testing.T) {

	var runner day10.Runner

	inputs := []byte("89010123\n78121874\n87430965\n96549874\n45678903\n32019012\n01329801\n10456732")

	res, err := runner.Part1(inputs)

	if err != nil {
		t.Fatalf("There shouldn't be error: %s\n", err.Error())
	}

	if res != 36 {
		t.Fatalf("Should return 36 instead %d\n", res)
	}
}

func Test3(t *testing.T) {

	var runner day10.Runner

	inputs := []byte("9990999\n9991999\n9992999\n6543456\n7111117\n8111118\n9111199")

	res, err := runner.Part1(inputs)

	if err != nil {
		t.Fatalf("There shouldn't be error: %s\n", err.Error())
	}

	if res != 2 {
		t.Fatalf("Should return 2 instead %d\n", res)
	}
}

func Test4(t *testing.T) {

	var runner day10.Runner

	inputs := []byte("1190119\n1111198\n1112117\n6543456\n7651987\n8761111\n9871111")

	res, err := runner.Part1(inputs)

	if err != nil {
		t.Fatalf("There shouldn't be error: %s\n", err.Error())
	}

	if res != 4 {
		t.Fatalf("Should return 4 instead %d\n", res)
	}
}
