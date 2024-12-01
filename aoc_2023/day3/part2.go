package day3

import (
	"strconv"
)

func checkGearAround(x, y, maxY, maxX int, buffer []byte, lines [][]byte) (Position, int) {

	var SPECIAL_CHAR byte = 42

	size := len(buffer)

	for i, _ := range buffer {

		neighbors := []Position{
			{y: y - 1, x: x + i},
			{y: y + 1, x: x + i},
		}

		if i == 0 {
			neighbors = append(neighbors,
				Position{y: y - 1, x: x - 1},
				Position{y: y + 1, x: x - 1},
				Position{y: y, x: x - 1},
			)
		}
		if i == size-1 {
			neighbors = append(neighbors,
				Position{y: y - 1, x: x + 1 + i},
				Position{y: y + 1, x: x + 1 + i},
				Position{y: y, x: x + 1 + i},
			)
		}

		for _, position := range neighbors {

			if position.isInRange(maxY, maxX) && lines[position.y][position.x] == SPECIAL_CHAR {
				r, _ := strconv.Atoi(string(buffer))
				return position, r
			}

		}

	}

	return Position{}, 0

}

func findNumbersNextGears(lines [][]byte) (res int) {

	var MAX_LINES = len(lines)
	var MAX_LINE = len(lines[0])

	var gears = make(map[Position][]int)

	for y := 0; y < MAX_LINES; y++ {
		for x := 0; x < MAX_LINE; x++ {

			letter := lines[y][x]

			if letter >= 48 && letter <= 57 {
				buffer, index := findNumber(lines[y][x:])

				position, number := checkGearAround(x, y, MAX_LINES, MAX_LINE, buffer, lines)

				if number != 0 {
					gears[position] = append(gears[position], number)
				}

				x += index

			}
		}
	}

	for _, v := range gears {
		if len(v) == 2 {
			res += v[0] * v[1]
		}
	}

	return res
}
