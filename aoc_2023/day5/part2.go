package day5

import (
	"fmt"
	"strings"
)

type SeedPair struct {
	start  int
	end    int
	length int
}

func initSeedsPair(firstLine string) (seeds []*SeedPair) {

	numbers := findNumbers(firstLine)

	for i := 0; i < len(numbers); i += 2 {

		seeds = append(seeds, &SeedPair{
			start:  numbers[i],
			end:    numbers[i] + numbers[i+1],
			length: numbers[i+1],
		})

	}

	return seeds
}

func totalSeeds(seeds []*SeedPair) (res int) {
	for _, value := range seeds {
		res += value.length
	}

	return res
}

func findLowestLocationWithPair(data string) int {

	lines := strings.Split(data, "\n\n")

	seeds := initSeedsPair(lines[0])

	fmt.Println(totalSeeds(seeds))

	return 0

}
