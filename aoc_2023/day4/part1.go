package day4

import (
	"bytes"
	"math"
	"regexp"
	"strconv"
)

func computeWinningNumbers(line []byte) (res float64) {

	var SEPARATOR_INDEX = 40
	var converted int
	var guess = make(map[int]int)

	regex := regexp.MustCompile(`\d+`)

	start := bytes.IndexByte(line, ':') + 2

	foundedNumbers := regex.FindAllSubmatch(line[SEPARATOR_INDEX:], -1)
	winningNumbers := regex.FindAllSubmatch(line[start:SEPARATOR_INDEX], -1)

	for _, a := range foundedNumbers {
		converted, _ = strconv.Atoi(string(a[0]))
		guess[converted] = converted
	}

	for _, a := range winningNumbers {
		converted, _ = strconv.Atoi(string(a[0]))
		_, exist := guess[converted]

		if exist {
			res++
		}
	}

	if res > 0 {
		return math.Pow(2, res-1)
	}

	return 0

}
