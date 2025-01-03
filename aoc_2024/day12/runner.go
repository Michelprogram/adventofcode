package day12

import (
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

func (a Area) findConnectedAndTracked(point utils.Point[rune]) (map[utils.Point[rune]]struct{}, []utils.Point[rune]) {

	connected := make(map[utils.Point[rune]]struct{})
	tracks := make([]utils.Point[rune], 0)

	for _, p := range FindAdjacents(point) {
		if _, exist := a.Plants[p]; exist {
			connected[p] = struct{}{}
			tracks = append(tracks, p)
		}
	}

	return connected, tracks

}

func (g Garden) CheckCorners(point utils.Point[rune], connected map[utils.Point[rune]]struct{}, side *int) {

	pointExists := func(p utils.Point[rune]) bool {
		_, exists := connected[p]
		return exists
	}

	corners := map[string]struct {
		check      func(p utils.Point[rune]) bool
		cornerFunc func(p utils.Point[rune]) utils.Point[rune]
	}{
		"topRight": {
			check: func(p utils.Point[rune]) bool {
				return pointExists(utils.Point[rune]{X: p.X, Y: p.Y - 1, Value: p.Value}) &&
					pointExists(utils.Point[rune]{X: p.X + 1, Y: p.Y, Value: p.Value})
			},
			cornerFunc: g.topRight,
		},
		"bottomRight": {
			check: func(p utils.Point[rune]) bool {
				return pointExists(utils.Point[rune]{X: p.X + 1, Y: p.Y, Value: p.Value}) &&
					pointExists(utils.Point[rune]{X: p.X, Y: p.Y + 1, Value: p.Value})
			},
			cornerFunc: g.bottomRight,
		},
		"bottomLeft": {
			check: func(p utils.Point[rune]) bool {
				return pointExists(utils.Point[rune]{X: p.X, Y: p.Y + 1, Value: p.Value}) &&
					pointExists(utils.Point[rune]{X: p.X - 1, Y: p.Y, Value: p.Value})
			},
			cornerFunc: g.bottomLeft,
		},
		"topLeft": {
			check: func(p utils.Point[rune]) bool {
				return pointExists(utils.Point[rune]{X: p.X - 1, Y: p.Y, Value: p.Value}) &&
					pointExists(utils.Point[rune]{X: p.X, Y: p.Y - 1, Value: p.Value})
			},
			cornerFunc: g.topLeft,
		},
	}

	for _, config := range corners {
		if config.check(point) && config.cornerFunc(point).Value != point.Value {
			*side++
			if len(connected) == 2 {
				break
			}
		}
	}
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

func (g Garden) topRight(point utils.Point[rune]) utils.Point[rune] {
	return g.Map[point.Y-1][point.X+1]
}

func (g Garden) bottomRight(point utils.Point[rune]) utils.Point[rune] {
	return g.Map[point.Y+1][point.X+1]
}
func (g Garden) bottomLeft(point utils.Point[rune]) utils.Point[rune] {
	return g.Map[point.Y+1][point.X-1]
}

func (g Garden) topLeft(point utils.Point[rune]) utils.Point[rune] {
	return g.Map[point.Y-1][point.X-1]
}

func (g Garden) ComputeSides(area Area) (side int) {

	if len(area.Plants) == 1 {
		return 4
	}

	for key := range area.Plants {

		connected, tracks := area.findConnectedAndTracked(key)

		switch len(tracks) {
		case 1:
			side += 2
		case 2:

			p1, p2 := tracks[0], tracks[1]

			//Straight doesn't count
			if p1.X == p2.X || p1.Y == p2.Y {
				continue
			}

			side++

			g.CheckCorners(key, connected, &side)
		case 3, 4:
			g.CheckCorners(key, connected, &side)

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

	res := 0

	garden := NewGarden(data)

	visited := make(map[utils.Point[rune]]struct{})

	for _, plant := range garden.Plants {

		area, _ := garden.FindRegionArea(plant, visited)

		if area != nil {
			res += garden.ComputeSides(*area) * len(area.Plants)
		}

	}

	return res, nil
}

func init() {
	registry.RegisterChallenge(12, Runner{})
}
