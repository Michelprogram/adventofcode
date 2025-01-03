package day6

import (
	"bytes"
	"errors"
	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
)

type Runner struct{}

type Grid struct {
	Inputs      [][]byte
	Cursor      [2]int
	Orientation [4][2]int
	MaxX        int
	MaxY        int
}

func NewGrid(inputs [][]byte) *Grid {
	return &Grid{
		Inputs: inputs,
		Cursor: [2]int{},
		MaxY:   len(inputs),
		MaxX:   len(inputs[0]),
		Orientation: [4][2]int{
			//x,y
			{0, -1},
			{1, 0},
			{0, 1},
			{-1, 0},
		},
	}
}

var _ utils.Challenge = (*Runner)(nil)

func (d Runner) ParseInputs(data []byte) ([][]byte, error) {
	return bytes.Split(data, []byte("\n")), nil
}

func (g *Grid) findCursor() error {

	for y, row := range g.Inputs {
		for x, letter := range row {
			if letter == byte('^') {
				g.Cursor = [2]int{x, y}
				g.Inputs[y][x] = '.'
				return nil
			}
		}
	}

	return errors.New("cursor not found")
}

func (g Grid) isInGrid(x, y int) bool {
	return x >= 0 && x < g.MaxX && y >= 0 && y < g.MaxY
}

func (g Grid) collision(x, y int, i *int) bool {

	if g.Inputs[y][x] == byte('#') {
		if *i+1 == len(g.Orientation) {
			*i = 0
		} else {
			*i++
		}
		return true
	}

	return false
}

func (g Grid) walk() int {

	i := 0

	visited := make(map[[2]int]struct{})

	for {
		// Calculate the next position
		nextX := g.Cursor[0] + g.Orientation[i][0]
		nextY := g.Cursor[1] + g.Orientation[i][1]

		if !g.isInGrid(nextX, nextY) {
			break
		}

		if g.collision(nextX, nextY, &i) {
			continue
		}

		visited[[2]int{nextX, nextY}] = struct{}{}

		g.Cursor = [2]int{nextX, nextY}
	}

	return len(visited)

}

func (d Runner) Part1(data []byte) (any, error) {

	inputs, _ := d.ParseInputs(data)

	grid := NewGrid(inputs)

	err := grid.findCursor()

	if err != nil {
		return nil, err
	}

	return grid.walk(), nil
}

func (d Runner) Part2(data []byte) (any, error) {

	var res int

	return res, nil
}

func init() {
	registry.RegisterChallenge(6, Runner{})
}
