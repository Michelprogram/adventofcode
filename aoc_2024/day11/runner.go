package day11

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	registry "github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
)

type Runner struct{}

var _ utils.Challenge = (*Runner)(nil)

func (d Runner) ParseInput(data []byte) ([]int, error) {
	res := make([]int, 0)

	for _, number := range strings.Fields(string(data)) {

		num, err := strconv.Atoi(string(number))

		if err != nil {
			return nil, err
		}

		res = append(res, num)
	}

	return res, nil
}

func (d Runner) Worker(data []int, iteration int) int {

	for i := 0; i < iteration; i++ {

		tmp := make([]int, 0, len(data)*2)

		for _, number := range data {

			if number == 0 {
				tmp = append(tmp, 1)
			} else if size := utils.NumberOfDigits(number); size%2 == 0 {

				divisor := int(math.Pow10(size / 2))

				leftPart, rightPart := number/divisor, number%divisor
				tmp = append(tmp, leftPart, rightPart)
			} else {
				tmp = append(tmp, number*2024)
			}

		}

		data = tmp
	}

	return len(data)
}

func (d Runner) WorkerWithMemoization(data []int, iteration int) int {

	memo := make(map[string]int)

	var runner func(number, level int) int
	runner = func(number, level int) int {
		if level == iteration {
			return 1
		}

		key := fmt.Sprintf("%d_%d", number, level)
		if count, exists := memo[key]; exists {
			return count
		}

		var result int
		if number == 0 {
			result = runner(1, level+1)
		} else if size := utils.NumberOfDigits(number); size%2 == 0 {
			divisor := int(math.Pow10(size / 2))
			leftPart := number / divisor
			rightPart := number % divisor
			result = runner(leftPart, level+1) + runner(rightPart, level+1)
		} else {
			result = runner(number*2024, level+1)
		}
		memo[key] = result
		return result
	}

	totalStones := 0
	for _, num := range data {
		totalStones += runner(num, 0)
	}

	return totalStones

}

func (d Runner) Part1(data []byte) (any, error) {

	input, err := d.ParseInput(data)

	if err != nil {
		return nil, err
	}

	return d.Worker(input, 25), nil
}

func (d Runner) Part2(data []byte) (any, error) {
	input, err := d.ParseInput(data)

	if err != nil {
		return nil, err
	}

	return d.WorkerWithMemoization(input, 75), nil
}

func init() {
	registry.RegisterChallenge(11, Runner{})
}
