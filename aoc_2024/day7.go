package aoc2024

import (
	"bytes"
	"strconv"

	"github.com/michelprogram/adventofcode/utils"
)

type Day7 struct{}

var _ utils.Challenge = (*Day7)(nil)

func (d Day7) ParseInputs(data []byte) (map[int][]int, error) {

	res := make(map[int][]int)

	for _, line := range bytes.Split(data, []byte("\n")) {

		colonIndex := bytes.Index(line, []byte(":"))

		if colonIndex == -1 {
			continue
		}

		key, err := strconv.Atoi(string(line[:colonIndex]))

		if err != nil {
			return nil, err
		}

		numbers := bytes.Fields(line[colonIndex+2:])

		res[key] = make([]int, len(numbers))

		for i, number := range numbers {

			n, err := strconv.Atoi(string(number))

			if err != nil {
				return nil, err
			}

			res[key][i] = n

		}

	}

	return res, nil

}

func (d Day7) generateCombinations(numbers []int, index int, current []string, results *[][]string) {

	if index >= len(numbers)-1 {
		*results = append(*results, append([]string{}, current...))
		return
	}

	if index == len(numbers)-1 {
		*results = append(*results, current)
		return
	}

	d.generateCombinations(numbers, index+1, append(current, "+"), results)
	d.generateCombinations(numbers, index+1, append(current, "*"), results)
}

func (d Day7) isEquationResolvable(expected int, numbers []int, possibilities [][]string) bool {

	for _, possibility := range possibilities {

		res := numbers[0]

		for i := 1; i < len(numbers); i++ {
			switch possibility[i-1] {
			case "+":
				res += numbers[i]
			case "*":
				res *= numbers[i]
			}
		}

		if res == expected {
			return true
		}
	}

	return false

}

func (d Day7) Part1(data []byte) (any, error) {

	var res int

	inputs, err := d.ParseInputs(data)

	if err != nil {
		return nil, err
	}

	for key, values := range inputs {

		possibilities := [][]string{}

		d.generateCombinations(values, 0, []string{}, &possibilities)

		if d.isEquationResolvable(key, values, possibilities) {
			res += key
		}

	}

	return res, nil
}

func (d Day7) Part2(data []byte) (any, error) {

	var res int

	return res, nil
}