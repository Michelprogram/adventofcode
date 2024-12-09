package aoc2024

import (
	"strconv"
	"strings"

	"github.com/michelprogram/adventofcode/utils"
)

type Day5 struct{}

var _ utils.Challenge = (*Day5)(nil)

func (d Day5) ParseInputs(data []byte) (map[string]map[string]struct{}, [][]string, error) {

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

func (d Day5) Part1(data []byte) (any, error) {
	var res int

	//inputs := []byte("47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47")

	hash, updates, err := d.ParseInputs(data)

	if err != nil {
		return nil, err
	}

	for _, numbers := range updates {
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

		if isValid {

			n, err := strconv.Atoi(numbers[len(numbers)/2])

			if err != nil {
				return nil, err
			}

			res += n

		}
	}

	return res, nil
}

func (d Day5) Part2(data []byte) (any, error) {

	var res int

	return res, nil
}
