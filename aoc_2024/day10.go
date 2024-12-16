package aoc2024

import (
	"github.com/michelprogram/adventofcode/utils"
	"strconv"
	"strings"
)

type Day10 struct{}

var _ utils.Challenge = (*Day10)(nil)

type Trail struct {
	Map     [][]int
	Visited map[Point]struct{}
	Heads   []Point
	MaxY    int
	MaxX    int
}

func NewTrail(data [][]int) *Trail {
	return &Trail{
		Map:     data,
		Visited: make(map[Point]struct{}),
		Heads:   make([]Point, 0),
		MaxY:    len(data),
		MaxX:    len(data[0]),
	}
}

func (t Trail) FindAdjacent(point Point, expected int) []Point {
	points := make([]Point, 0, 4)

	cross := []Point{
		{
			point.X + 1, point.Y,
		},
		{
			point.X, point.Y - 1,
		},
		{
			point.X - 1, point.Y,
		},
		{
			point.X, point.Y + 1,
		},
	}

	for _, p := range cross {
		if !utils.IsOutOfBound(p.X, p.Y, t.MaxX, t.MaxY) && t.Map[p.Y][p.X] == expected {
			points = append(points, p)
		}
	}

	return points
}

func (t *Trail) SetHeads() {

	for i := 0; i < len(t.Map); i++ {
		for j := 0; j < len(t.Map[0]); j++ {
			if t.Map[i][j] == 0 {
				t.Heads = append(t.Heads, Point{i, j})
			}
		}
	}
}

func (t Trail) FindPath(head Point, start int) int {

	if start == 9 {
		return 1
	}

	points := t.FindAdjacent(head, start)

	for _, p := range points {
		t.FindPath(p, start+1)
	}
}

func (t *Trail) Scores() int {

	for _, head := range t.Heads {
		t.FindPath(head, 1)
	}

	return len(t.Visited)

	/*	for i := 0; i < 9; i++ {
			temp := make(map[Point]struct{})
			for point, _ := range t.Points {
				for _, adjacent := range t.FindAdjacent(point, i+1) {
					temp[adjacent] = struct{}{}
				}
			}
			t.Points = temp
		}

		return len(t.Points)*/
}

func (d Day10) ParseInputs(data []byte) ([][]int, error) {
	inputs := strings.Split(string(data), "\n")
	var result [][]int

	for i := 0; i < len(inputs)-1; i++ {
		nums := strings.Split(inputs[i], "")
		tempo := make([]int, 0, len(nums))
		for n := 0; n < len(nums); n++ {
			num, err := strconv.Atoi(nums[n])
			if err != nil {
				return nil, err
			}
			tempo = append(tempo, num)
		}
		result = append(result, tempo)

	}

	return result, nil
}

func (d Day10) Part1(data []byte) (any, error) {

	data = []byte("89010123\n78121874\n87430965\n96549874\n45678903\n32019012\n01329801\n10456732\n")

	inputs, err := d.ParseInputs(data)

	if err != nil {
		return nil, err
	}

	trail := NewTrail(inputs)

	trail.SetHeads()

	return trail.Scores(), nil
}

func (d Day10) Part2(data []byte) (any, error) {

	return nil, nil
}
