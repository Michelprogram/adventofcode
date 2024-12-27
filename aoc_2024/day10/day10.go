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

type Node struct {
	utils.Point
	Value int
}

func (t Node) String() string {
	return fmt.Sprintf("Point at {x:%d;y:%d} with value %d", t.X, t.Y, t.Value)
}

type Graph struct {
	Nodes   map[utils.Point]*Node
	Started []*Node
}

func NewGraph(data []byte) (*Graph, error) {

	lines := strings.Fields(string(data))

	graph := &Graph{
		Nodes:   make(map[utils.Point]*Node),
		Started: make([]*Node, 0),
	}

	for y, line := range lines {
		for x, number := range strings.Split(line, "") {

			num, err := strconv.Atoi(number)

			if err != nil {
				return nil, err
			}

			node := &Node{utils.Point{X: x, Y: y}, num}

			if num == 0 {
				graph.Started = append(graph.Started, node)
			}

			graph.Nodes[node.Point] = node
		}
	}

	return graph, nil
}

func (t Graph) FindPath(node Node) int {

	var worker func(nestedLevel int, node Node)
	visited := make(map[utils.Point]struct{})
	founded := make(map[utils.Point]struct{})

	worker = func(nestedLevel int, node Node) {

		if nestedLevel == 9 && node.Value == 9 {
			founded[node.Point] = struct{}{}
			return
		}

		visited[node.Point] = struct{}{}

		expected := node.Value + 1
		adjacents := []utils.Point{
			{node.X, node.Y - 1},
			{node.X, node.Y + 1},
			{node.X + 1, node.Y},
			{node.X - 1, node.Y},
		}

		for _, adjacent := range adjacents {

			next, isExist := t.Nodes[adjacent]
			_, alreadyVisited := visited[adjacent]
			if isExist && next.Value == expected && !alreadyVisited {
				worker(nestedLevel+1, *next)
			}

		}
		delete(visited, node.Point)
	}

	worker(0, node)

	return len(founded)

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

	return nil, nil
}
