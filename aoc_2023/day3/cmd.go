package day3

import (
	"bytes"
	"github.com/michelprogram/adventofcode2023/utils"
)

type Day3 struct {
}

var _ utils.Code = (*Day3)(nil)

func (d Day3) Execute(dataSet []byte) (any, error) {

	/*	data := [][]byte{
			[]byte("467..114.."),
			[]byte("...*......"),
			[]byte("..35..633."),
			[]byte("......#..."),
			[]byte("617*......"),
			[]byte(".....+.58."),
			[]byte("..592....."),
			[]byte("......755."),
			[]byte("...$.*...."),
			[]byte(".664.598.."),
		}

		return findNumbersNextGears(data), nil*/

	//return findNumbersNextToSymbol(data), nil

	lines := bytes.Split(dataSet, []byte("\n"))

	/*	return findNumbersNextGears(lines[:len(lines)-1]), nil
	 */
	return findNumbersNextToSymbol(lines[:len(lines)-1]), nil

}
