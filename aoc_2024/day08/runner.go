package day8

import (
	"bytes"
	"fmt"
	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
	"log"
)

type Runner struct{}

var _ utils.Challenge = (*Runner)(nil)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {

	return fmt.Sprintf("X:%d Y:%d", p.X, p.Y)
}

func (p Point) createAntinodes(point Point) Point {

	var x, y int

	diffX := utils.Abs(p.X - point.X)
	diffY := utils.Abs(p.Y - point.Y)

	//Is under
	if point.Y > p.Y {
		y = p.Y - diffY
	} else {
		y = p.Y + diffY
	}

	if point.X > p.X {
		x = p.X - diffX
	} else {
		x = p.X + diffX
	}

	return Point{
		X: x,
		Y: y,
	}
}

type Map struct {
	Antennas  map[byte][]Point
	Antinodes map[Point]struct{}
	Grid      [][]byte
	MaxX      int
	MaxY      int
	Counter   int
}

func NewMap(data [][]byte) *Map {
	return &Map{
		Antennas:  make(map[byte][]Point),
		Antinodes: make(map[Point]struct{}),
		Grid:      data,
		MaxX:      len(data[0]),
		MaxY:      len(data),
		Counter:   0,
	}
}

func (g *Map) setAntennas() {

	for y, row := range g.Grid {
		for x, val := range row {
			if val != byte('.') {

				if _, ok := g.Antennas[val]; !ok {
					g.Antennas[val] = make([]Point, 0)
				}
				g.Antennas[val] = append(g.Antennas[val], Point{x, y})
			}
		}
	}
}

func (g Map) isOutOfBound(point Point) bool {
	return point.X < 0 || point.X >= g.MaxX || point.Y >= g.MaxY || point.Y < 0
}

func (g *Map) generateAntiNode() {

	for _, points := range g.Antennas {
		for i := 0; i < len(points); i++ {
			for j := 0; j < len(points); j++ {
				if i != j {
					antinode := points[i].createAntinodes(points[j])
					if !g.isOutOfBound(antinode) {
						g.Antinodes[antinode] = struct{}{}
						if g.Grid[antinode.Y][antinode.X] == byte('.') {
							g.Grid[antinode.Y][antinode.X] = byte('#')
						}
					}
				}

			}
		}
	}
}

func (g *Map) generateAntiNodeResonant() {

	for _, points := range g.Antennas {
		for i := 0; i < len(points); i++ {
			for j := 0; j < len(points); j++ {
				if i != j {
					antinode := points[i].createAntinodes(points[j])
					last := points[i]
					for !g.isOutOfBound(antinode) {

						g.Antinodes[antinode] = struct{}{}

						if g.Grid[antinode.Y][antinode.X] == byte('.') {
							g.Grid[antinode.Y][antinode.X] = byte('#')
						}

						tmp := antinode

						antinode = antinode.createAntinodes(last)

						last = tmp
					}
				}
			}
		}
	}
}

func (g Map) String() string {
	var res = "\n"

	for _, col := range g.Grid {
		res += fmt.Sprintf("%s\n", col)
	}

	return res
}

func (d Runner) ParseInputs(data []byte) ([][]byte, error) {

	inputs := bytes.Split(data, []byte("\n"))

	return inputs[:len(inputs)-1], nil
}

func (d Runner) Part1(data []byte) (any, error) {

	inputs, _ := d.ParseInputs(data)

	grid := NewMap(inputs)

	grid.setAntennas()

	grid.generateAntiNode()

	return len(grid.Antinodes), nil
}

func (d Runner) Part2(data []byte) (any, error) {

	inputs, _ := d.ParseInputs(data)

	grid := NewMap(inputs)

	grid.setAntennas()

	grid.generateAntiNodeResonant()

	log.Println(len(grid.Antinodes), grid.Antinodes)

	return len(grid.Antinodes), nil
}

func init() {
	registry.RegisterChallenge(8, Runner{})
}
