package day8

import (
	"github.com/michelprogram/adventofcode2023/utils"
)

type Day8 struct {
}

var _ utils.Code = (*Day8)(nil)

func (d Day8) Execute(dataSet []byte) (any, error) {

	//data := "RL\n\nAAA = (BBB, CCC)\nBBB = (DDD, EEE)\nCCC = (ZZZ, GGG)\nDDD = (DDD, DDD)\nEEE = (EEE, EEE)\nGGG = (GGG, GGG)\nZZZ = (ZZZ, ZZZ)"

	//data := "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)"

	size := len(dataSet)

	return camelDirection(string(dataSet[:size-1])), nil

	//return camelDirection(string(data)), nil

}
