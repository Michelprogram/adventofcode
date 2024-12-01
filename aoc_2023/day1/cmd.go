package day1

import (
	"github.com/michelprogram/adventofcode2023/utils"
	"strings"
)

type Day1 struct {
}

var _ utils.Code = (*Day1)(nil)

func (D Day1) Execute(dataSet []byte) (any, error) {

	res := 0

	words := strings.Split(string(dataSet), "\n")

	for _, word := range words {
		res += findNumbers(word)
	}

	return res, nil
}
