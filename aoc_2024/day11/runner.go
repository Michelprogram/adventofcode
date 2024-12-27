package day11

import (
	"log"
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

	cache := make(map[int][]int)

	cache[0] = []int{1}

	for i := 0; i < iteration; i++ {
		log.Println(i)
		tmp := make([]int, 0, len(data)*2)

		for _, number := range data {

			res, exist := cache[number]

			if exist {
				tmp = append(tmp, res...)
			} else if size := utils.NumberOfDigits(number); size%2 == 0 {

				divisor := int(math.Pow10(size / 2))

				leftPart, rightPart := number/divisor, number%divisor
				cache[number] = []int{leftPart, rightPart}
				tmp = append(tmp, leftPart, rightPart)
			} else {
				tmp = append(tmp, number*2024)
				cache[number] = []int{number * 2024}
			}

		}

		data = tmp
	}

	log.Println(data)

	return len(data)
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

	return d.Worker(input, 75), nil
}

func init() {
	registry.RegisterChallenge(11, Runner{})
}
