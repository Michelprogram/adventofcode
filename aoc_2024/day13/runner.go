package day13

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
)

type Runner struct{}

type Button struct {
	X int
	Y int
}

type Machine struct {
	BtnA  Button
	BtnB  Button
	Prize Button
}

func NewMachine(params []string) (*Machine, error) {

	integers := make([]int, 6)

	for i, param := range params {
		num, err := strconv.Atoi(param)

		if err != nil {
			return nil, err
		}

		integers[i] = num
	}

	return &Machine{
		BtnA:  Button{X: integers[0], Y: integers[1]},
		BtnB:  Button{X: integers[2], Y: integers[3]},
		Prize: Button{X: integers[4], Y: integers[5]},
	}, nil

}

func (m Machine) LowestToken() int {

	lowest := 1_000_000

	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {
			px := m.BtnA.X*i + m.BtnB.X*j
			py := m.BtnA.Y*i + m.BtnB.Y*j

			if px == m.Prize.X && py == m.Prize.Y {
				lowest = utils.Min(lowest, 3*i+j)
			}
		}
	}

	if lowest < 1_000_000 {
		return lowest
	}

	return 0

}

var _ utils.Challenge = (*Runner)(nil)

func (d Runner) ParseInput(data []byte) ([]Machine, error) {
	blocks := strings.Split(string(data), "\n\n")
	machines := make([]Machine, 0, len(blocks))

	r, _ := regexp.Compile("\\d+")

	for _, block := range blocks {

		founded := r.FindAllString(block, -1)

		temp, err := NewMachine(founded)

		if err != nil {
			return nil, err
		}

		machines = append(machines, *temp)

	}

	return machines, nil
}

func (d Runner) Part1(data []byte) (any, error) {

	res := 0

	machines, err := d.ParseInput(data)

	if err != nil {
		return nil, err
	}

	for _, machine := range machines {
		res += machine.LowestToken()
	}

	return res, nil
}

func (d Runner) Part2(data []byte) (any, error) {
	// TODO: Implement Part 2 logic here
	return nil, nil
}

func init() {
	registry.RegisterChallenge(13, Runner{})
}
