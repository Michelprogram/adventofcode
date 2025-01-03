package day5

import (
	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
	"log"
	"strconv"
	"strings"
)

type Runner struct{}

var _ utils.Challenge = (*Runner)(nil)

func (d Runner) ParseInputs(data []byte) (map[string]map[string]struct{}, [][]string, error) {

	pageOrdering := make(map[string]map[string]struct{})

	inputs := strings.Split(string(data), "\n\n")

	for _, rules := range strings.Split(inputs[0], "\n") {
		rule := strings.Split(rules, "|")

		left, right := rule[0], rule[1]

		if _, exists := pageOrdering[left]; !exists {
			pageOrdering[left] = make(map[string]struct{})
		}

		pageOrdering[left][right] = struct{}{}
	}

	lines := strings.Split(inputs[1], "\n")

	updates := make([][]string, len(lines)-1)

	for i := 0; i < len(lines)-1; i++ {

		numbers := strings.Split(lines[i], ",")

		updates[i] = numbers
	}

	return pageOrdering, updates, nil
}

func (d Runner) isValidSequence(numbers []string, hash map[string]map[string]struct{}) bool {
	seen := make(map[string]struct{})
	isValid := true

	for _, number := range numbers {
		for k := range hash[number] {
			if _, exists := seen[k]; exists {
				isValid = false
				break
			}
		}
		if !isValid {
			break
		}
		seen[number] = struct{}{}
	}

	return isValid
}

func (d Runner) Part1(data []byte) (any, error) {
	var res int

	hash, updates, err := d.ParseInputs(data)

	if err != nil {
		return nil, err
	}

	for _, numbers := range updates {

		if d.isValidSequence(numbers, hash) {

			n, err := strconv.Atoi(numbers[len(numbers)/2])

			if err != nil {
				return nil, err
			}

			res += n

		}
	}

	return res, nil
}

func (d Runner) orderSequence(numbers []string, hash map[string]map[string]struct{}) string {

	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if _, ok := hash[numbers[j+1]][numbers[j]]; ok {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	//97,13,75,29,47
	return numbers[len(numbers)/2]

}

func (d Runner) Part2(data []byte) (any, error) {

	var res int

	hash, updates, err := d.ParseInputs(data)

	if err != nil {
		return nil, err
	}

	for _, numbers := range updates {

		if !d.isValidSequence(numbers, hash) {

			log.Println(numbers)

			middle := d.orderSequence(numbers, hash)

			number, err := strconv.Atoi(middle)

			if err != nil {
				return nil, err
			}

			res += number

		}

	}

	return res, nil
}

func init() {
	registry.RegisterChallenge(5, Runner{})
}
