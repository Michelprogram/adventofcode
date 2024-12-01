package day2

import (
	"bytes"
	"github.com/michelprogram/adventofcode2023/utils"
)

type Day2 struct {
}

var _ utils.Code = (*Day2)(nil)

func (d Day2) Execute(dataSet []byte) (any, error) {

	res := 0

	words := bytes.Split(dataSet, []byte("\n"))

	for _, line := range words {
		//res += correctBag(line)
		res += findMaxColorPerBag(line)
	}

	return res, nil
}
