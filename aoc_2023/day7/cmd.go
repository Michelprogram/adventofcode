package day7

import (
	"github.com/michelprogram/adventofcode2023/utils"
)

type Day7 struct {
}

var _ utils.Code = (*Day7)(nil)

func (d Day7) Execute(dataSet []byte) (any, error) {

	//data := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"

	//data = "25J7T 46\n22893 698\nJJJJJ 700\nJJJ8J 300"

	size := len(dataSet)

	//return cardsRank(string(dataSet[:size-1])), nil

	return cardsRankJoker(string(dataSet[:size-1])), nil
}
