package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
)

func init() {
	registry.RegisterChallenge(10, Runner{})
}

type Runner struct{}

var _ utils.Challenge = (*Runner)(nil)

type Storage interface {
	Add(utils.Point[any])
}

type MapStorage struct {
	Data map[utils.Point[any]]struct{}
}

func (m *MapStorage) Add(point utils.Point[any]) {
	m.Data[point] = struct{}{}
}

type ArrayStorage struct {
	Data []utils.Point[any]
}

func (a *ArrayStorage) Add(point utils.Point[any]) {
	a.Data = append(a.Data, point)
}

type Node struct {
	utils.Point[any]
	Value int
}

func (t Node) String() string {
	return fmt.Sprintf("Point at {x:%d;y:%d} with value %d", t.X, t.Y, t.Value)
}

type Graph struct {
	Nodes   map[utils.Point[any]]*Node
	Started []*Node
}

func NewGraph(data []byte) (*Graph, error) {

	lines := strings.Fields(string(data))

	graph := &Graph{
		Nodes:   make(map[utils.Point[any]]*Node),
		Started: make([]*Node, 0),
	}

	for y, line := range lines {
		for x, number := range strings.Split(line, "") {

			num, err := strconv.Atoi(number)

			if err != nil {
				return nil, err
			}

			node := &Node{utils.Point[any]{X: x, Y: y}, num}

			if num == 0 {
				graph.Started = append(graph.Started, node)
			}

			graph.Nodes[node.Point] = node
		}
	}

	return graph, nil
}

func (t Graph) worker(nestedLevel int, node Node, visited map[utils.Point[any]]struct{}, storage Storage) {
	if nestedLevel == 9 && node.Value == 9 {
		storage.Add(node.Point)
		//		founded[node.Point] = struct{}{}
		return
	}

	visited[node.Point] = struct{}{}

	expected := node.Value + 1
	adjacents := []utils.Point[any]{
		{X: node.X, Y: node.Y - 1},
		{X: node.X, Y: node.Y + 1},
		{X: node.X + 1, Y: node.Y},
		{X: node.X - 1, Y: node.Y},
	}

	for _, adjacent := range adjacents {

		next, isExist := t.Nodes[adjacent]
		_, alreadyVisited := visited[adjacent]
		if isExist && next.Value == expected && !alreadyVisited {
			t.worker(nestedLevel+1, *next, visited, storage)
		}

	}
	delete(visited, node.Point)
}

func (t Graph) FindPath(node Node) int {

	visited := make(map[utils.Point[any]]struct{})
	founded := MapStorage{make(map[utils.Point[any]]struct{})}

	t.worker(0, node, visited, &founded)

	return len(founded.Data)
}

func (t Graph) FindPathV2(node Node) int {

	visited := make(map[utils.Point[any]]struct{})
	founded := ArrayStorage{make([]utils.Point[any], 0)}

	t.worker(0, node, visited, &founded)

	return len(founded.Data)

}

func (d Runner) Part1(data []byte) (any, error) {

	res := 0

	graph, err := NewGraph(data)

	if err != nil {
		return nil, err
	}

	for _, start := range graph.Started {
		res += graph.FindPath(*start)
	}

	return res, nil

}

func (d Runner) Part2(data []byte) (any, error) {

	res := 0

	graph, err := NewGraph(data)

	if err != nil {
		return nil, err
	}

	for _, start := range graph.Started {
		res += graph.FindPathV2(*start)
	}

	return res, nil

}
