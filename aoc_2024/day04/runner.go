package day4

import (
	"bytes"
	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
	"strings"
)

type Day4 struct {
	Inputs   []string
	MaxY     int
	MaxX     int
	Word     string
	Reverse  string
	WordSize int
	Counter  int
	Inputs2  [][]byte
}

type Runner struct{}

var _ utils.Challenge = (*Runner)(nil)

func (d Runner) ParseInputs(data []byte) []string {

	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func NewDay4(inputs []string) *Day4 {

	//Micro optimisation don't need X
	return &Day4{
		Inputs:   inputs,
		Word:     "MAS",
		Reverse:  "SAM",
		MaxY:     len(inputs),
		MaxX:     len(inputs[0]),
		WordSize: len("MAS"),
		Counter:  0,
	}
}

func (d Runner) Part1(data []byte) (any, error) {

	day := NewDay4(d.ParseInputs(data))

	for y, line := range day.Inputs {
		for x, letter := range line {
			if letter == rune('X') {
				day.findWord(x, y)
			}
		}
	}

	return day.Counter, nil

}

func (d *Day4) findWord(x, y int) {

	d.FindHorizontal(x, y)
	d.FindVertical(x, y)
	d.FindCross(x, y)

}

func (d *Day4) FindHorizontal(x, y int) {

	//x to Right
	if x+d.WordSize < d.MaxX && d.Inputs[y][x+1:x+d.WordSize+1] == d.Word {
		d.Counter++
	}

	//x to Left
	if x-d.WordSize >= 0 && d.Inputs[y][x-(d.WordSize):x] == d.Reverse {
		d.Counter++
	}
}

func (d *Day4) FindVertical(x, y int) {

	var i int

	//y to Top
	if y-d.WordSize >= 0 {

		for i = 1; i <= d.WordSize; i++ {
			if d.Inputs[y-i][x] != d.Word[i-1] {
				break
			}
		}

		if i == d.WordSize+1 {
			d.Counter++
		}
	}

	//y to Bot
	if y+d.WordSize < d.MaxY {

		for i = 1; i <= d.WordSize; i++ {
			if d.Inputs[y+i][x] != d.Word[i-1] {
				break
			}
		}

		if i == d.WordSize+1 {
			d.Counter++
		}

	}
}

func (d *Day4) FindCross(x, y int) {

	var i int

	//Top left corner
	if x-d.WordSize >= 0 && y-d.WordSize >= 0 {
		for i = 1; i <= d.WordSize; i++ {
			if d.Inputs[y-i][x-i] != d.Word[i-1] {
				break
			}
		}

		if i == d.WordSize+1 {
			d.Counter++
		}
	}

	//Top right corner
	if x+d.WordSize < d.MaxX && y-d.WordSize >= 0 {

		for i = 1; i <= d.WordSize; i++ {
			if d.Inputs[y-i][x+i] != d.Word[i-1] {
				break
			}
		}

		if i == d.WordSize+1 {
			d.Counter++
		}
	}

	//Bot left corner
	if x-d.WordSize >= 0 && y+d.WordSize < d.MaxY {
		for i = 1; i <= d.WordSize; i++ {
			if d.Inputs[y+i][x-i] != d.Word[i-1] {
				break
			}
		}

		if i == d.WordSize+1 {

			d.Counter++
		}
	}

	//Bot right corner
	if x+d.WordSize < d.MaxX && y+d.WordSize < d.MaxY {
		for i = 1; i <= d.WordSize; i++ {
			if d.Inputs[y+i][x+i] != d.Word[i-1] {
				break
			}
		}

		if i == d.WordSize+1 {
			d.Counter++
		}
	}

}

func (d Day4) isXMas(x, y int) bool {
	matrix := [][]int{
		{-1, -1, 1, 1},
		{-1, 1, 1, -1},
	}

	for _, coord := range matrix {
		res := []byte{d.Inputs2[y-coord[0]][x-coord[1]], d.Inputs2[y-coord[2]][x-coord[3]]}

		if bytes.Compare(res, []byte("MS")) != 0 && bytes.Compare(res, []byte("SM")) != 0 {
			return false
		}
	}

	return true
}

func (d Runner) Part2(data []byte) (any, error) {

	var res int

	var day Day4

	inputs := bytes.Split(data, []byte("\n"))

	day.Inputs2 = inputs

	for y := 1; y < len(inputs)-2; y++ {
		for x := 1; x < len(inputs[0])-1; x++ {
			if inputs[y][x] == byte('A') && day.isXMas(x, y) {
				res++
			}
		}
	}

	return res, nil
}
func init() {
	registry.RegisterChallenge(4, Runner{})
}
