package day1

import (
	"strconv"
)

func findNumbers(text string) int {

	size := len(text)
	i := 0
	res := make([]byte, 2)

	for i = 0; i < size; i++ {
		if text[i] >= 48 && text[i] <= 57 {
			res[0] = text[i]
			break
		}
	}

	for i = size - 1; i >= 0; i-- {
		if text[i] >= 48 && text[i] <= 57 {
			res[1] = text[i]
			break
		}
	}

	r, _ := strconv.Atoi(string(res))

	return r
}
