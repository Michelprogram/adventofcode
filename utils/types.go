package utils

type Code interface {
	Execute(data []byte, part, day int) (any, error)
}

type Challenge interface {
	Part1([]byte) (any, error)
	Part2([]byte) (any, error)
}

type Point[T any] struct {
	X     int
	Y     int
	Value T
}
