package day3

import (
	"strconv"
)

type Position struct {
	x int
	y int
}

func (p Position) isInRange(maxY, maxX int) bool {
	return p.y > 0 && p.x > 0 && p.y < maxY && p.x < maxX
}

func findNumber(line []byte) (numberBuffer []byte, index int) {

	for index < len(line) {
		char := line[index]

		if char >= 48 && char <= 57 {
			numberBuffer = append(numberBuffer, char)
			index++
		} else {
			break
		}
	}

	return numberBuffer, index

}

func checkAround(x, y, maxY, maxX int, buffer []byte, lines [][]byte) int {

	var SPECIAL_CHARS = map[byte]bool{
		33: true,
		34: true,
		35: true,
		36: true,
		37: true,
		38: true,
		39: true,
		40: true,
		41: true,
		42: true,
		43: true,
		44: true,
		45: true,
		47: true,
		58: true,
		59: true,
		60: true,
		61: true,
		62: true,
		63: true,
		64: true,
	}

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

			if position.isInRange(maxY, maxX) && SPECIAL_CHARS[lines[position.y][position.x]] {
				r, _ := strconv.Atoi(string(buffer))
				return r
			}

		}

	}

	return 0

}

func findNumbersNextToSymbol(lines [][]byte) (res int) {

	var MAX_LINES = len(lines)
	var MAX_LINE = len(lines[0])

	for y := 0; y < MAX_LINES; y++ {
		for x := 0; x < MAX_LINE; x++ {

			letter := lines[y][x]

			if letter >= 48 && letter <= 57 {
				buffer, index := findNumber(lines[y][x:])

				output := checkAround(x, y, MAX_LINES, MAX_LINE, buffer, lines)

				if output != 0 {
					res += output
				}

				x += index

			}
		}
	}

	return res

}
