package day12

import (
	"log"
	"strings"

	"github.com/Michelprogram/data-structures/stack"
	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
)

type Runner struct{}

type Garden struct {
	MaxY      int
	MaxX      int
	Validator map[utils.Point[rune]]struct{}
	Map       [][]utils.Point[rune]
	Plants    []utils.Point[rune]
}

type Area struct {
	Plants   map[utils.Point[rune]]struct{}
	Letter   rune
	MaxPlant utils.Point[rune]
	MinPlant utils.Point[rune]
}

func (a Area) IsExist(x, y int) bool {
	_, exist := a.Plants[utils.Point[rune]{X: x, Y: y, Value: a.Letter}]

	return exist
}

var _ utils.Challenge = (*Runner)(nil)

func FindAdjacents(p utils.Point[rune]) []utils.Point[rune] {
	return []utils.Point[rune]{
		{X: p.X, Y: p.Y - 1, Value: p.Value},
		{X: p.X + 1, Y: p.Y, Value: p.Value},
		{X: p.X, Y: p.Y + 1, Value: p.Value},
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
func (g Garden) FindRegionArea(point utils.Point[rune], visited map[utils.Point[rune]]struct{}) (*Area, error) {

	if _, ok := visited[point]; ok {
		return nil, nil
	}

	maxPoint := utils.Point[rune]{X: -1, Y: -1}

	minPoint := utils.Point[rune]{X: g.MaxX + 1, Y: g.MaxY + 1}

	s := stack.NewStack[utils.Point[rune]](1_000)

	res := make(map[utils.Point[rune]]struct{})
	_, _ = s.Push(point)

	for !s.IsEmpty() {

		item, _ := s.Pop()

		if _, exist := visited[item]; exist {
			continue
		}

		if item.X <= minPoint.X && item.Y <= minPoint.Y {
			minPoint = item
		}

		if item.X >= maxPoint.X && item.Y >= maxPoint.Y {
			maxPoint = item
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

	return &Area{
		Plants:   res,
		Letter:   point.Value,
		MaxPlant: maxPoint,
		MinPlant: minPoint,
	}, nil
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

	if len(area) == 1 {
		return 4
	}

	for key := range area {

		connected := make([]utils.Point[rune], 0)

		for _, p := range FindAdjacents(key) {
			if _, exist := area[p]; exist {
				connected = append(connected, p)
			}
		}

		switch len(connected) {
		case 1:
			side += 2
		case 2:

			p1, p2 := connected[0], connected[1]

			diffX := utils.Abs(p1.X - p2.X)
			diffY := utils.Abs(p1.Y - p2.Y)

			if diffX == 0 || diffY == 0 {
				continue
			}

			if p1.Y == key.Y-1 && p1.X == key.X && p2.X == key.X+1 && p2.Y == key.Y {
				if g.Map[key.Y-1][key.X+1].Value != key.Value {
					if key.Value == rune('C') {
						log.Println("1")
					}
					side++
				}
			} else if p1.X == key.X+1 && p1.Y == key.Y && p2.Y == key.Y+1 && p2.X == key.X {
				if g.Map[key.Y+1][key.X+1].Value != key.Value {
					side++
				}
			} else if p1.Y == key.Y+1 && p1.X == key.X && p2.X == key.X-1 && p2.Y == key.Y {
				if g.Map[key.Y+1][key.X-1].Value != key.Value {
					if key.Value == rune('C') {
						log.Println("2")
					}
					side++
				}
			} else if p1.X == key.X-1 && p1.Y == key.Y && p2.Y == key.Y-1 && p2.X == key.X {
				if g.Map[key.Y-1][key.X-1].Value != key.Value {
					side++
				}
			}

			side++

		}
	}

	return side
}

func (d Runner) Part1(data []byte) (any, error) {

	computed := 0

	garden := NewGarden(data)

	visited := make(map[utils.Point[rune]]struct{})

	for _, plant := range garden.Plants {

		area, _ := garden.FindRegionArea(plant, visited)

		if area != nil {
			computed += len(area.Plants) * garden.ComputePerimeter(area.Plants)
		}

	}

	return computed, nil
}

func (d Runner) Part2(data []byte) (any, error) {

	data = []byte("EEEEE\nEXXXX\nEEEEE\nEXXXX\nEEEEE")

	garden := NewGarden(data)

	visited := make(map[utils.Point[rune]]struct{})

	for _, plant := range garden.Plants {

		area, _ := garden.FindRegionArea(plant, visited)

		if area != nil {
			_ = garden.ComputeSides(area.Plants)
		}

	}

	return nil, nil
}

func init() {
	registry.RegisterChallenge(12, Runner{})
}
