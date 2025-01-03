//go:build js && wasm

package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"syscall/js"
)

func main() {
	js.Global().Set("dfs", dfsWrapper())
	<-make(chan struct{})
}

func dfsWrapper() js.Func {
	jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			return "Unable to get document object"
		}

		garden := NewGarden(args[0].String(), jsDoc)

		visited := make(map[Point]struct{})

		for _, plant := range garden.Plants {

			_ = garden.dfs(plant, visited)

		}

		return nil
	})

	return jsonfunc
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
}

func FindAdjacents(p Point) []Point {
	return []Point{
		{X: p.X, Y: p.Y - 1, Value: p.Value},
		{X: p.X + 1, Y: p.Y, Value: p.Value},
		{X: p.X, Y: p.Y + 1, Value: p.Value},
		{X: p.X - 1, Y: p.Y, Value: p.Value},
	}
}

func NewGarden(data string, jsdoc js.Value) *Garden {

	lines := strings.Split(data, "\n")

	res := make([][]Point, len(lines))

	garden := &Garden{
		Plants:    make([]Point, 0),
		Validator: make(map[Point]struct{}),
		JsDoc:     jsdoc,
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

func (g Garden) setColorHtml(point Point, color string) error {
	selector := fmt.Sprintf("span[data-x='%d'][data-y='%d']", point.X, point.Y)
	element := g.JsDoc.Call("querySelector", selector)
	if !element.IsNull() {
		element.Get("style").Set("color", color)
	}

	return nil
}

func (g Garden) dfs(point Point, visited map[Point]struct{}) error {

	if _, ok := visited[point]; ok {
		return nil
	}

	color, _ := g.pickRandomColor()

	s := NewStack[Point](1_000)

	_, _ = s.Push(point)

	for !s.IsEmpty() {

		item, _ := s.Pop()

		if _, exist := visited[item]; exist {
			continue
		}

		visited[item] = struct{}{}

		go func() {
			g.setColorHtml(item, color)
		}()

		adjacents := FindAdjacents(item)

		for _, adjacent := range adjacents {
			_, exist := g.Validator[adjacent]
			_, alreadyVisited := visited[adjacent]

			if exist && !alreadyVisited {
				_, _ = s.Push(adjacent)
			}

		}
	}

	return nil
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
