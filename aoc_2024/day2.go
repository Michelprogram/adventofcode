package aoc2024

import (
	"bytes"
	"github.com/michelprogram/adventofcode/utils"
	"strconv"
)

type Day2 struct{}

var _ utils.Challenge = (*Day2)(nil)

func (d Day2) ParseInputs(data []byte) ([][]int, error) {

	lines := bytes.Split(data, []byte("\n"))

	result := make([][]int, len(lines)-1)

	for i := 0; i < len(lines)-1; i++ {
		numbers := bytes.Split(lines[i], []byte(" "))

		nested := make([]int, len(numbers))

		for j, number := range numbers {
			n, err := strconv.Atoi(string(number))

			if err != nil {
				return nil, err
			}

			nested[j] = n
		}

		result[i] = nested
	}

	return result, nil

}

func (d Day2) Part1(data []byte) (any, error) {

	var counter = 0

	inputs, err := d.ParseInputs(data)

	if err != nil {
		return counter, err
	}

	for _, input := range inputs {

		min := input[0]
		max := input[len(input)-1]

		if (max > min && d.Backward(input)) || d.Forward(input) {
			counter++
		}

	}

	return counter, nil

}

func (d Day2) Backward(input []int) bool {
	for i := len(input) - 1; i > 0; i-- {
		if input[i] <= input[i-1] || input[i]-input[i-1] > 3 {
			return false
		}
	}

	return true
}

func (d Day2) Forward(input []int) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i+1] >= input[i] || input[i]-input[i+1] > 3 {
			return false
		}
	}

	return true
}

func (d Day2) Part2(data []byte) (any, error) {

	return nil, nil
}
