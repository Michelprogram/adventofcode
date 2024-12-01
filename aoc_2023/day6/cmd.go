package day6

import (
	"github.com/michelprogram/adventofcode2023/utils"
)

type Day6 struct {
}

var _ utils.Code = (*Day6)(nil)

func (d Day6) Execute(dataSet []byte) (any, error) {

	data := [][]byte{
		[]byte("Time:        56     71     79     99"),
		[]byte("Distance:   334   1135   1350   2430"),
	}

	return boatRaceWithoutSpace(data), nil

}
