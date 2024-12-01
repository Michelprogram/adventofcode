package day2

import (
	"strconv"
)

func findNumber(line []byte) (number, index int) {

	var numberBuffer []byte

	index = 0
	char := line[index]

	for char >= 48 && char <= 57 {
		numberBuffer = append(numberBuffer, char)
		index++
		char = line[index]
	}

	number, _ = strconv.Atoi(string(numberBuffer))

	return number, index

}

func getId(line []byte) (res, index int) {
	return findNumber(line[5:])
}

func isCorrect(colors map[string]int) bool {
	return colors["green"] <= 13 && colors["blue"] <= 14 && colors["red"] <= 12
}

func correctBag(line []byte) int {

	var colors = make(map[string]int)
	var tempo int
	var index int

	size := len(line)

	if size == 0 {
		return 0
	}

	id, start := getId(line)

	start += 7

	for i := start; i < size; i++ {
		letter := line[i]

		//Semicolon
		if letter == 59 {
			if isCorrect(colors) {
				colors = make(map[string]int)
			} else {
				return 0
			}

		}

		//Is number
		if letter >= 48 && letter <= 57 {
			tempo, index = findNumber(line[i:])
			i += index
		} else if letter == 98 {
			colors["blue"] = tempo
			i += 3

		} else if letter == 103 {
			colors["green"] = tempo
			i += 4

		} else if letter == 114 {
			colors["red"] += tempo
			i += 2

		}
	}

	if isCorrect(colors) {
		return id
	}

	return 0

}
