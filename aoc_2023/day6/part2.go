package day6

import (
	"fmt"
	"regexp"
	"strconv"
)

func formatData2(data [][]byte) (races Race) {

	var time string

	var distance string

	regex := regexp.MustCompile(`\d+`)

	times := regex.FindAllSubmatch(data[0], -1)
	distances := regex.FindAllSubmatch(data[1], -1)

	for i := 0; i < len(times); i++ {

		time += string(times[i][0])
		distance += string(distances[i][0])

	}

	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(distance)

	return Race{
		time:     t,
		distance: d,
	}
}

func boatRaceWithoutSpace(data [][]byte) int {

	race := formatData2(data)

	var times []int

	res := 1

	fmt.Println(race)

	for holdButtonTime := 0; holdButtonTime <= race.time; holdButtonTime++ {
		remainingTime := race.time - holdButtonTime
		distanceCovered := holdButtonTime * remainingTime

		if distanceCovered > race.distance {
			times = append(times, holdButtonTime)
		}
	}

	res *= len(times)
	times = []int{}

	return res
}
