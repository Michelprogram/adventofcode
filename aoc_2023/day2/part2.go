package day2

func checkMax(colors map[string]int, save map[string]int) {
	for k, v := range colors {
		if save[k] < v {
			save[k] = v
		}
	}
}

func findMaxColorPerBag(line []byte) int {

	var tempo int
	var index int

	colors := make(map[string]int)

	save := map[string]int{
		"green": 0,
		"blue":  0,
		"red":   0,
	}

	size := len(line)

	if size == 0 {
		return 0
	}

	_, start := getId(line)

	start += 7

	for i := start; i < size; i++ {
		letter := line[i]

		//Semicolon
		if letter == 59 {
			checkMax(colors, save)
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
			colors["red"] = tempo
			i += 2

		}
	}

	checkMax(colors, save)

	return save["red"] * save["green"] * save["blue"]

}
