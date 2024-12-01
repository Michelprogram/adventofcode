package day5

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Seed struct {
	value int
	flag  bool
}

func findLowestValue(seeds []*Seed) int {

	var min = math.MaxInt

	for _, seed := range seeds {
		if seed.value < min {
			min = seed.value
		}
	}

	return min
}

func findNumbers(line string) []int {
	regex := regexp.MustCompile(`\d+`)

	numbers := regex.FindAllString(line, -1)

	res := make([]int, len(numbers))

	for i, number := range numbers {
		r, _ := strconv.Atoi(number)
		res[i] = r
	}

	return res
}

func resetFlag(seeds *[]*Seed) {

	for _, seed := range *seeds {
		seed.flag = false
	}

}

func initSeeds(firstLine string) (seeds []*Seed) {

	for _, seed := range findNumbers(firstLine) {
		seeds = append(seeds, &Seed{
			value: seed,
			flag:  false,
		})
	}

	return seeds
}

func findLowestLocation(data string) int {

	lines := strings.Split(data, "\n\n")

	seeds := initSeeds(lines[0])

	for _, line := range lines[1:] {
		cuted := strings.Split(line, "\n")
		//fmt.Printf("%#v\n", cuted)
		for _, numbersOnLine := range cuted[1:] {
			numbers := findNumbers(numbersOnLine)
			destination, source, length := numbers[0], numbers[1], numbers[2]

			for _, seed := range seeds {
				if seed.value >= source && seed.value <= source+length && !seed.flag {
					//fmt.Printf("Computed : %d + (%d - %d) = %d\n", destination, seed.value, source, destination+(seed.value-source))
					seed.value = destination + (seed.value - source)
					seed.flag = true
					//fmt.Printf("Destination %d, source %d, length %d, seed %d\n", destination, source, length, seed.value)
				}
			}

		}
		resetFlag(&seeds)
	}

	return findLowestValue(seeds)

}
