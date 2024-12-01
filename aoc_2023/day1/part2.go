package day1

import (
	"strconv"
)

func findNumbersWithLetters(text string) int {

	hash := map[string]byte{
		"one":   49,
		"two":   50,
		"three": 51,
		"four":  52,
		"five":  53,
		"six":   54,
		"seven": 55,
		"eight": 56,
		"nine":  57,
		"eno":   49,
		"owt":   50,
		"eerht": 51,
		"ruof":  52,
		"evif":  53,
		"xis":   54,
		"neves": 55,
		"thgie": 56,
		"enin":  57,
	}

	var temp []byte
	var letter byte

	size := len(text)
	i := 0
	res := make([]byte, 2)

	for i = 0; i < size; i++ {
		letter = text[i]
		if letter >= 48 && letter <= 57 {
			res[0] = letter
		} else {
			temp = append(temp, letter)
		}

		if len(temp) >= 3 {
			for j := 0; j < len(temp); j++ {
				cuted := temp[j:]
				v, exist := hash[string(cuted)]

				if exist {
					res[0] = v
				}

			}
		}

		if res[0] != 0 {
			break
		}

	}

	temp = []byte{}

	for i = size - 1; i >= 0; i-- {
		letter = text[i]
		if letter >= 48 && letter <= 57 {
			res[1] = letter
			break
		} else {
			temp = append(temp, letter)
		}

		if len(temp) >= 3 {
			for j := 0; j < len(temp); j++ {
				cuted := temp[j:]
				v, exist := hash[string(cuted)]

				if exist {
					res[1] = v
				}

			}
		}

		if res[1] != 0 {
			break
		}
	}

	r, _ := strconv.Atoi(string(res))

	return r
}
