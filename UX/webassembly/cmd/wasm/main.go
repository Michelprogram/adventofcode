//go:build js && wasm

package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"syscall/js"
	"time"
)

func main() {
	js.Global().Set("dfs", js.FuncOf(dfsWrapper))
	<-make(chan struct{})

}

func dfsWrapper(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return "Invalid no of arguments passed"
	}

	start := time.Now()

	garden := NewGarden(args[0].String())

	visited := make(map[Point]struct{})

	res := make([][]Point, 0)

	for _, plant := range garden.Plants {

		points := garden.dfs(plant, visited)

		if points != nil {
			res = append(res, points)
		}

	}

	log.Println(len(res), time.Since(start))
	j, _ := json.Marshal(res)

	return string(j)
}

type Point struct {
	X, Y  int
	Value rune
}

type Garden struct {
	Validator map[Point]struct{}
	Map       [][]Point
	Plants    []Point
	JsDoc     js.Value
	Canva     js.Value
}

func FindAdjacents(p Point) []Point {
	return []Point{
		{X: p.X, Y: p.Y - 1, Value: p.Value},
		{X: p.X + 1, Y: p.Y, Value: p.Value},
		{X: p.X, Y: p.Y + 1, Value: p.Value},
		{X: p.X - 1, Y: p.Y, Value: p.Value},
	}
}

func NewGarden(data string) *Garden {

	lines := strings.Split(data, "\n")

	res := make([][]Point, len(lines))

	garden := &Garden{
		Plants:    make([]Point, 0),
		Validator: make(map[Point]struct{}),
	}

	for y, line := range lines {
		temp := make([]Point, len(line))
		for x, char := range line {
			pt := Point{Y: y, X: x, Value: char}
			temp[x] = pt
			garden.Plants = append(garden.Plants, pt)
			garden.Validator[pt] = struct{}{}
		}

		res[y] = temp
	}

	garden.Map = res

	return garden
}

func (g Garden) pickRandomColor() (string, error) {
	bytes := make([]byte, 3)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return "#" + hex.EncodeToString(bytes), nil
}

func (g Garden) dfs(point Point, visited map[Point]struct{}) []Point {

	res := make([]Point, 0)

	if _, ok := visited[point]; ok {
		return nil
	}

	s := NewStack[Point](1_000)

	_, _ = s.Push(point)

	res = append(res, point)

	for !s.IsEmpty() {

		item, _ := s.Pop()

		if _, exist := visited[item]; exist {
			continue
		}

		visited[item] = struct{}{}

		res = append(res, item)

		adjacents := FindAdjacents(item)

		for _, adjacent := range adjacents {
			_, exist := g.Validator[adjacent]
			_, alreadyVisited := visited[adjacent]

			if exist && !alreadyVisited {
				_, _ = s.Push(adjacent)
			}

		}
	}

	return res
}

type Stack[T any] struct {
	arr  []T
	top  int
	size int
}

func NewStack[T any](size int) *Stack[T] {

	stack := Stack[T]{
		arr:  make([]T, size),
		top:  0,
		size: size,
	}

	return &stack

}

//Add an element to the top of a stack

func (s *Stack[T]) Push(element T) (T, error) {

	if s.IsFull() {

		return nilItem[T](), fmt.Errorf("The stack is full !")
	}

	s.arr[s.top] = element

	s.top++

	return element, nil
}

// Remove an element from the top of a stack
func (s *Stack[T]) Pop() (T, error) {

	if s.IsEmpty() {

		return nilItem[T](), fmt.Errorf("The stack is empty !")
	}

	item := s.arr[s.top-1]

	s.top--

	return item, nil
}

//Check if the stack is empty

func (s Stack[T]) IsEmpty() bool {
	return s.top == 0
}

// Check if the stack is full
func (s Stack[T]) IsFull() bool {
	return s.top == s.size
}

// Get the value of the top element without removing it
func (s Stack[T]) Peek() T {
	return s.arr[s.top-1]
}

// Show what we have inside Stack
func (s Stack[T]) String() string {
	var res string

	for index, element := range s.arr {
		res += fmt.Sprintf("Index : %d Value : %v\n", index, element)
	}

	return res
}

func nilItem[T any]() T {
	var item T
	return item
}
