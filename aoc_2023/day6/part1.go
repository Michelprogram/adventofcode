package day6

import (
	"regexp"
	"strconv"
)

type Race struct {
	time     int
	distance int
}

func formatData(data [][]byte) (races []Race) {
	regex := regexp.MustCompile(`\d+`)

	times := regex.FindAllSubmatch(data[0], -1)
	distances := regex.FindAllSubmatch(data[1], -1)

	for i := 0; i < len(times); i++ {

		time, _ := strconv.Atoi(string(times[i][0]))
		distance, _ := strconv.Atoi(string(distances[i][0]))

		races = append(races, Race{
			time, distance,
		})
	}

	return races
}

func boatRace(data [][]byte) int {

	races := formatData(data)

	var times []int

	res := 1

	for _, race := range races {
		for holdButtonTime := 0; holdButtonTime <= race.time; holdButtonTime++ {
			remainingTime := race.time - holdButtonTime
			distanceCovered := holdButtonTime * remainingTime

			if distanceCovered > race.distance {
				times = append(times, holdButtonTime)
			}
		}
		res *= len(times)
		times = []int{}
	}

	return res
}
