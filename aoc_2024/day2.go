package aoc2024

import (
	"bytes"
	"strconv"

	"github.com/michelprogram/adventofcode/utils"
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

	inputs, err := d.ParseInputs(data)

	if err != nil {
		return nil, err
	}

	counter := 0

	for _, numbers := range inputs {
		if d.isSafe(numbers) {
			counter++
		}
	}

	return counter, nil
}

func (d Day2) isSafe(numbers []int) bool {

	isIncreasing := numbers[0] < numbers[1]

	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]

		if diff < -3 || diff > 3 {
			return false
		}

		if isIncreasing && diff <= 0 {
			return false
		}
		if !isIncreasing && diff >= 0 {
			return false
		}
	}
	return true
}

func (d Day2) Part2(data []byte) (any, error) {

	inputs, err := d.ParseInputs(data)

	if err != nil {
		return nil, err
	}

	counter := 0

	for _, numbers := range inputs {
		if d.isSafeWithDampener(numbers) {
			counter++
		}
	}

	return counter, nil

}

func (d Day2) isSafeWithDampener(numbers []int) bool {
	if d.isSafe(numbers) {
		return true
	}

	for i := 0; i < len(numbers); i++ {
		modified := append([]int{}, numbers[:i]...)
		modified = append(modified, numbers[i+1:]...)

		if d.isSafe(modified) {
			return true
		}
	}

	return false
}
