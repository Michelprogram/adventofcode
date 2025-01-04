package day14

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
)

type Runner struct{}

type Robot struct {
	Position utils.Point[struct{}]
	Velocity utils.Point[struct{}]
}

func NewRobot(params string) (*Robot, error) {

	r, _ := regexp.Compile("-?\\d+")
	numbers := make([]int, 4)

	for i, number := range r.FindAllString(params, 4) {
		n, err := strconv.Atoi(number)

		if err != nil {
			return nil, err
		}

		numbers[i] = n

	}

	return &Robot{
		Position: utils.Point[struct{}]{
			X: numbers[0],
			Y: numbers[1],
		},
		Velocity: utils.Point[struct{}]{
			X: numbers[2],
			Y: numbers[3],
		},
	}, nil

}

func (r *Robot) Moove(maxX, maxY int) {
	if r.Position.X+r.Velocity.X >= maxX {
		r.Position.X = (r.Position.X + r.Velocity.X) - maxX
	} else if r.Position.X+r.Velocity.X < 0 {
		r.Position.X = maxX + (r.Position.X + r.Velocity.X)
	} else {
		r.Position.X += r.Velocity.X
	}
	if r.Position.Y+r.Velocity.Y >= maxY {
		r.Position.Y = (r.Position.Y + r.Velocity.Y) - maxY
	} else if r.Position.Y+r.Velocity.Y < 0 {
		r.Position.Y = maxY + (r.Position.Y + r.Velocity.Y)
	} else {
		r.Position.Y += r.Velocity.Y
	}
}

var _ utils.Challenge = (*Runner)(nil)

func (d Runner) ParseInput(data []byte) ([]*Robot, error) {
	lines := strings.Split(string(data), "\n")
	res := make([]*Robot, 0, len(lines))

	for _, line := range lines {

		if line == "" {
			continue
		}

		robot, err := NewRobot(line)

		if err != nil {
			return nil, err
		}

		res = append(res, robot)
	}

	return res, nil
}

func (d Runner) Display(maxX, maxY int, robots []*Robot, output io.Writer) {
	grid := make([][]int, maxY)

	for i := range grid {
		grid[i] = make([]int, maxX)
	}

	for _, robot := range robots {
		grid[robot.Position.Y][robot.Position.X]++
	}

	for _, row := range grid {
		for _, count := range row {
			if count == 0 {
				output.Write([]byte("."))
			} else {
				output.Write([]byte("#"))
			}
		}
		output.Write([]byte("\n"))
	}
	output.Write([]byte("\n\n"))
}

func (d Runner) Iteration(iteration int, robots []*Robot, maxX, maxY int) {
	for i := 0; i < iteration; i++ {
		for _, robot := range robots {
			robot.Moove(maxX, maxY)
		}
	}
}

func (d Runner) IsBetweenPoint(point, min, max utils.Point[struct{}]) bool {

	return point.X >= min.X && point.Y >= min.Y && point.X <= max.X && point.Y <= max.Y

}

func (d Runner) FindRobotInQuadrant(maxX, maxY int, robots []*Robot) (res int) {

	type sector struct {
		Min utils.Point[struct{}]
		Max utils.Point[struct{}]
	}

	halfX, halfY := maxX/2, maxY/2

	counters := make([]int, 4)

	sectors := [4]sector{
		{
			Min: utils.Point[struct{}]{X: 0, Y: 0},
			Max: utils.Point[struct{}]{X: halfX - 1, Y: halfY - 1},
		},
		{
			Min: utils.Point[struct{}]{X: halfX + 1, Y: 0},
			Max: utils.Point[struct{}]{X: maxX - 1, Y: halfY - 1},
		},
		{
			Min: utils.Point[struct{}]{X: 0, Y: halfY + 1},
			Max: utils.Point[struct{}]{X: halfX - 1, Y: maxY - 1},
		},
		{
			Min: utils.Point[struct{}]{X: halfX + 1, Y: halfY + 1},
			Max: utils.Point[struct{}]{X: maxX - 1, Y: maxY - 1},
		},
	}

	for _, robot := range robots {
		for i, sector := range sectors {
			if d.IsBetweenPoint(robot.Position, sector.Min, sector.Max) {
				counters[i]++
				break
			}
		}
	}

	res = counters[0]

	for i := 1; i < 4; i++ {
		res *= counters[i]
	}

	return res
}

func (d Runner) FindAdjacentRecursivePointer(
	c *int,
	p utils.Point[struct{}],
	grid map[utils.Point[struct{}]]struct{},
	visited map[utils.Point[struct{}]]struct{},
) {
	visited[p] = struct{}{}

	adjacents := []utils.Point[struct{}]{
		{X: p.X, Y: p.Y - 1, Value: struct{}{}},
		{X: p.X + 1, Y: p.Y, Value: struct{}{}},
		{X: p.X, Y: p.Y + 1, Value: struct{}{}},
		{X: p.X - 1, Y: p.Y, Value: struct{}{}},
	}

	for _, adjacent := range adjacents {
		if _, alreadyVisited := visited[adjacent]; !alreadyVisited {
			if _, exist := grid[adjacent]; exist {
				*c++
				d.FindAdjacentRecursivePointer(c, adjacent, grid, visited)
			}
		}
	}
}

func (d Runner) FindAdjacentRecursive(
	p utils.Point[struct{}],
	grid map[utils.Point[struct{}]]struct{},
	visited map[utils.Point[struct{}]]struct{},
) int {
	visited[p] = struct{}{}

	count := 1

	adjacents := []utils.Point[struct{}]{
		{X: p.X, Y: p.Y - 1, Value: struct{}{}},
		{X: p.X + 1, Y: p.Y, Value: struct{}{}},
		{X: p.X, Y: p.Y + 1, Value: struct{}{}},
		{X: p.X - 1, Y: p.Y, Value: struct{}{}},
	}

	for _, adjacent := range adjacents {
		if _, alreadyVisited := visited[adjacent]; !alreadyVisited {
			if _, exist := grid[adjacent]; exist {
				count += d.FindAdjacentRecursive(adjacent, grid, visited)
			}
		}
	}

	return count
}

func (d Runner) Part1(data []byte) (any, error) {
	robots, err := d.ParseInput(data)

	if err != nil {
		return nil, err
	}

	maxX, maxY := 101, 103

	d.Iteration(100, robots, maxX, maxY)

	return d.FindRobotInQuadrant(maxX, maxY, robots), nil
}

func (d Runner) Part2(data []byte) (any, error) {

	robots, err := d.ParseInput(data)

	if err != nil {
		return nil, err
	}

	maxX, maxY := 101, 103

	for i := 0; i < 10_000; i++ {

		mappedRobot := make(map[utils.Point[struct{}]]struct{})

		for _, robot := range robots {
			mappedRobot[robot.Position] = struct{}{}
		}

		for _, robot := range robots {
			visited := make(map[utils.Point[struct{}]]struct{})
			if d.FindAdjacentRecursive(robot.Position, mappedRobot, visited) > 200 {
				return i, nil
			}
		}

		for _, robot := range robots {
			robot.Moove(maxX, maxY)
		}
	}

	return nil, nil
}

func init() {
	registry.RegisterChallenge(14, Runner{})
}
