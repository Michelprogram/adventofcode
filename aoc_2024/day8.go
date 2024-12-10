package aoc2024

import (
	"bytes"
	"github.com/michelprogram/adventofcode/utils"
	"log"
)

type Day8 struct{}

type Map struct {
	Antennas map[byte][][2]int
	Grid     [][]byte
}

func NewMap(data [][]byte) *Map {
	return &Map{
		Antennas: make(map[byte][][2]int),
		Grid:     data,
	}
}

func (g *Map) setAntennas() {

	for y, row := range g.Grid {
		for x, val := range row {
			if val != byte('.') {

				if _, ok := g.Antennas[val]; !ok {
					g.Antennas[val] = make([][2]int, 0)
				}
				g.Antennas[val] = append(g.Antennas[val], [2]int{x, y})
			}
		}
	}

}

func (g *Map) manhattanDist(x1, x2, y1, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

var _ utils.Challenge = (*Day8)(nil)

func (d Day8) ParseInputs(data []byte) ([][]byte, error) {

	inputs := bytes.Split(data, []byte("\n"))

	return inputs[:len(inputs)-1], nil
}

func (d Day8) Part1(data []byte) (any, error) {

	var res int

	data = []byte("..........\n...#......\n#.........\n....a.....\n........a.\n.....a....\n..#.......\n......#...\n..........\n..........")

	inputs, _ := d.ParseInputs(data)

	grid := NewMap(inputs)

	grid.setAntennas()

	coord := grid.Antennas[97]

	log.Println(grid.manhattanDist(coord[0][0], coord[0][1], coord[1][0], coord[1][1]))

	return res, nil
}

func (d Day8) Part2(data []byte) (any, error) {

	var res int

	return res, nil
}
