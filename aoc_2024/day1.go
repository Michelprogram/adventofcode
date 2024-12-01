package aoc2024

import (
	"bytes"
	"sort"
	"strconv"

	"github.com/michelprogram/adventofcode/utils"
)

type Day1 struct{}

var _ utils.Challenge = (*Day1)(nil)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseNumbers(data []byte) ([]int, []int, error) {

	lines := bytes.Split(data, []byte("\n"))

	max := len(lines)

	l1 := make([]int, max-1)
	l2 := make([]int, max-1)

	for i := 0; i < max-1; i++ {

		line := lines[i]

		space := bytes.Index(line, []byte("   "))

		if space == -1 {
			continue
		}

		res, err := strconv.Atoi(string(line[:space]))

		if err != nil {
			return nil, nil, err
		}
		l1[i] = res

		res, err = strconv.Atoi(string(line[space+3:]))

		if err != nil {
			return nil, nil, err
		}

		l2[i] = res
	}

	return l1, l2, nil

}

func (d Day1) Part1(data []byte) (any, error) {

	var res int

	l1, l2, err := parseNumbers(data)

	if err != nil {
		return nil, err
	}

	sort.Slice(l1, func(i, j int) bool {
		return l1[i] < l1[j]
	})

	sort.Slice(l2, func(i, j int) bool {
		return l2[i] < l2[j]
	})

	for i, n1 := range l1 {

		res += abs(n1 - l2[i])

	}

	return res, nil

}

func (d Day1) Part2(data []byte) (any, error) {

	var res int

	hash := make(map[int]int)

	l1, l2, err := parseNumbers(data)

	if err != nil {
		return nil, err
	}

	for _, number := range l2 {
		hash[number]++
	}

	for _, number := range l1 {
		res += (number * hash[number])
	}

	return res, nil
}
