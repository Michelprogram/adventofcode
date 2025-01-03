package day1

import (
	"bytes"
	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
	"sort"
	"strconv"
)

type Runner struct{}

var _ utils.Challenge = (*Runner)(nil)

func (d Runner) ParseInputs(data []byte) ([]int, []int, error) {

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

func (d Runner) Part1(data []byte) (any, error) {

	var res int

	l1, l2, err := d.ParseInputs(data)

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

		res += utils.Abs(n1 - l2[i])

	}

	return res, nil

}

func (d Runner) Part2(data []byte) (any, error) {

	var res int

	hash := make(map[int]int)

	l1, l2, err := d.ParseInputs(data)

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

func init() {
	registry.RegisterChallenge(1, Runner{})
}
