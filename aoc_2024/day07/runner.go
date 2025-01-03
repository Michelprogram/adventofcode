package day7

import (
	"bytes"
	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
	"strconv"
)

type Runner struct{}

var _ utils.Challenge = (*Runner)(nil)

func (d Runner) ParseInputs(data []byte) (map[int][]int, error) {

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

func (d Runner) generateCombinations(numbers []int, index int, current []string, results *[][]string) {

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
	d.generateCombinations(numbers, index+1, append(current, "||"), results)
}

func (d Runner) isEquationResolvable(expected int, numbers []int, possibilities [][]string) bool {

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

func (d Runner) Part1(data []byte) (any, error) {

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

func (d Runner) isEquationResolvable2(expected int, numbers []int, possibilities [][]string) bool {

	for _, possibility := range possibilities {
		res := numbers[0]

		for i := 1; i < len(numbers); i++ {
			switch possibility[i-1] {
			case "+":
				res += numbers[i]
			case "*":
				res *= numbers[i]
			case "||":
				res, _ = strconv.Atoi(strconv.Itoa(res) + strconv.Itoa(numbers[i]))
			}
		}
		if res == expected {
			return true
		}
	}

	return false

}

func (d Runner) Part2(data []byte) (any, error) {

	var res int

	//data = []byte("190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20\n")

	inputs, err := d.ParseInputs(data)

	if err != nil {
		return nil, err
	}

	for key, values := range inputs {

		possibilities := [][]string{}

		d.generateCombinations(values, 0, []string{}, &possibilities)

		if d.isEquationResolvable2(key, values, possibilities) {
			res += key
		}

	}

	return res, nil
}

func init() {
	registry.RegisterChallenge(7, Runner{})
}
