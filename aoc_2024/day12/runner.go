package day12

import (
	"github.com/Michelprogram/data-structures/stack"
	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
	"strings"
)

type Runner struct{}

type Garden struct {
	MaxY      int
	MaxX      int
	Validator map[utils.Point[rune]]struct{}
	Map       [][]utils.Point[rune]
	Plants    []utils.Point[rune]
}

var _ utils.Challenge = (*Runner)(nil)

func FindAdjacents(p utils.Point[rune]) []utils.Point[rune] {
	return []utils.Point[rune]{
		{X: p.X, Y: p.Y - 1, Value: p.Value},
		{X: p.X, Y: p.Y + 1, Value: p.Value},
		{X: p.X + 1, Y: p.Y, Value: p.Value},
		{X: p.X - 1, Y: p.Y, Value: p.Value},
	}
}

func NewGarden(data []byte) *Garden {

	lines := strings.Split(string(data), "\n")
	res := make([][]utils.Point[rune], len(lines))

	garden := &Garden{
		MaxX:      len(res[0]),
		MaxY:      len(res),
		Plants:    make([]utils.Point[rune], 0),
		Validator: make(map[utils.Point[rune]]struct{}),
	}

	for y, line := range lines {
		temp := make([]utils.Point[rune], len(line))
		for x, char := range line {

			pt := utils.Point[rune]{
				Y: y, X: x, Value: char,
			}

			temp[x] = pt
			garden.Plants = append(garden.Plants, pt)
			garden.Validator[pt] = struct{}{}
		}

		res[y] = temp
	}

	garden.Map = res

	return garden
}

// FindRegionArea Instead of doing recursion like day10 use stack
func (g Garden) FindRegionArea(point utils.Point[rune], visited map[utils.Point[rune]]struct{}) (map[utils.Point[rune]]struct{}, error) {

	if _, ok := visited[point]; ok {
		return nil, nil
	}

	s := stack.NewStack[utils.Point[rune]](1_000)

	res := make(map[utils.Point[rune]]struct{})
	_, _ = s.Push(point)

	for !s.IsEmpty() {

		item, _ := s.Pop()

		if _, exist := visited[item]; exist {
			continue
		}

		visited[item] = struct{}{}

		res[item] = struct{}{}

		adjacents := FindAdjacents(item)

		for _, adjacent := range adjacents {
			_, exist := g.Validator[adjacent]
			_, alreadyVisited := visited[adjacent]

			if exist && !alreadyVisited {
				_, _ = s.Push(adjacent)
			}

		}
	}

	if len(res) == 0 {
		return nil, nil
	}

	return res, nil
}

func (g Garden) ComputePerimeter(area map[utils.Point[rune]]struct{}) (res int) {

	for point := range area {

		base := 4

		adjacents := FindAdjacents(point)

		for _, adjacent := range adjacents {
			if _, exist := area[adjacent]; exist {
				base--
			}
		}

		res += base

	}

	return res

}

func (g Garden) ComputeSides(area map[utils.Point[rune]]struct{}) (side int) {

	visited := map[utils.Point[rune]]struct{}{}

	for point := range area {

	}

	return side
}

func (d Runner) Part1(data []byte) (any, error) {

	computed := 0

	garden := NewGarden(data)

	visited := make(map[utils.Point[rune]]struct{})

	for _, plant := range garden.Plants {

		res, _ := garden.FindRegionArea(plant, visited)

		if res != nil {
			computed += len(res) * garden.ComputePerimeter(res)
		}

	}
	return computed, nil
}

func (d Runner) Part2(data []byte) (any, error) {
	// TODO: Implement Part 2 logic here
	return nil, nil
}

func init() {
	registry.RegisterChallenge(12, Runner{})
}
