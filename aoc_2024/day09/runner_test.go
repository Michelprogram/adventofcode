package day09_test

import (
	"testing"

	"github.com/michelprogram/adventofcode/aoc_2024/day09"
)

type Provider struct {
	Inputs    []byte
	InputsInt []int
	CheckSum  int
}


func TestNewDisk(t *testing.T) {
	providers := []Provider{
		{
			Inputs:    []byte("2333133121414131402"),
			InputsInt: []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2},
			CheckSum:  1928,
		},
	}

	for _, provider := range providers {
		disk, err := day09.NewDisk(provider.Inputs)
		if err != nil {
			t.Fatalf("There shouldn't be error: %s\n", err.Error())
		}

		for i, container := range disk.Data {
			if container.Value != provider.InputsInt[i] {
				t.Fatalf("Values should be the same at index %d instead got %d and %d\n", i, container.Value, provider.InputsInt[i])
			}
		}
	}
}

func TestRunner_Part1(t *testing.T) {
	providers := []Provider{
		{
			Inputs:    []byte("2333133121414131402"),
			InputsInt: []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2},
			CheckSum:  1928,
		},
	}

	var runner day09.Runner

	for _, provider := range providers {
		checksum, err := runner.Part1(provider.Inputs)
		if err != nil {
			t.Fatalf("There shouldn't be error: %s\n", err.Error())
		}

		if checksum != provider.CheckSum {
			t.Fatalf("Checksum should be the same expected: %d get %d\n", provider.CheckSum, checksum)
		}

	}
}

func TestRunner_Part2(t *testing.T) {
	providers := []Provider{
		{
			Inputs:    []byte("2333133121414131402"),
			InputsInt: []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2},
			CheckSum:  2858,
		},
	}

	var runner day09.Runner

	for _, provider := range providers {
		res, err := runner.Part2(provider.Inputs)
		if err != nil {
			t.Fatalf("There shouldn't be error: %s\n", err.Error())
		}

		if res != provider.CheckSum {
			t.Fatalf("Result should be %d instead get %d\n", provider.CheckSum, res)
		}
	}
}
