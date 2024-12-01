package day4

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

func computeWinningScratchCards(line []byte, duplications *[]int, counter *int, j int) {

	//var SEPARATOR_INDEX = 40
	var SEPARATOR_INDEX = 23
	var BASE_SCRATCHCARD = 1
	var converted int
	var res int
	var loop = BASE_SCRATCHCARD
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

	size := len(*duplications)

	if size != 0 {
		loop = (*duplications)[0] + BASE_SCRATCHCARD
		*duplications = (*duplications)[1:]
		size = len(*duplications)
	}

	fmt.Printf("Card %d has %d matching and %d instances \n", j+1, res, loop)

	*counter += loop

	for i := 0; i < res; i++ {
		if i >= size {
			*duplications = append(*duplications, 1)
		} else {
			(*duplications)[i] += loop
		}
	}

}
