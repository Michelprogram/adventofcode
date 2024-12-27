package registry

import "github.com/michelprogram/adventofcode/utils"

var challenges = map[int]utils.Challenge{}

func RegisterChallenge(day int, challenge utils.Challenge) {
	challenges[day] = challenge
}

func GetChallenge(day int) (utils.Challenge, bool) {
	challenge, exists := challenges[day]
	return challenge, exists
}
