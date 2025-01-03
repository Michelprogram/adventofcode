package day3

import (
	"bytes"
	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
	"regexp"
	"strconv"
)

type Runner struct{}

var _ utils.Challenge = (*Runner)(nil)

func (d Runner) ParseInputs(data []byte) string {

	return string(bytes.TrimSpace(data))

}

func (d Runner) Part1(data []byte) (any, error) {

	var res int

	inputs := d.ParseInputs(data)

	regex := regexp.MustCompile("mul\\((?P<a>[0-9]*),(?P<b>[0-9]*)\\)")

	for _, groups := range regex.FindAllStringSubmatch(inputs, -1) {

		a, err := strconv.Atoi(groups[1])

		if err != nil {
			return nil, err
		}

		b, err := strconv.Atoi(groups[2])

		if err != nil {
			return nil, err
		}

		res += a * b
	}

	return res, nil

}

func (d Runner) Part2(data []byte) (any, error) {
	inputs := d.ParseInputs(data)

	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	regexDont := regexp.MustCompile(`don't`)

	var res int

	for i := 0; i < len(inputs); {
		dontMatch := regexDont.FindStringIndex(inputs[i:])
		mulMatch := regex.FindStringSubmatchIndex(inputs[i:])

		if len(dontMatch) > 0 && (len(mulMatch) == 0 || dontMatch[0] < mulMatch[0]) {
			i += dontMatch[1]
			for i < len(inputs)-4 && inputs[i:i+4] != "do()" {
				i++
			}
		} else if len(mulMatch) > 0 {
			a, err := strconv.Atoi(inputs[i+mulMatch[2] : i+mulMatch[3]])
			if err != nil {
				return nil, err
			}

			b, err := strconv.Atoi(inputs[i+mulMatch[4] : i+mulMatch[5]])
			if err != nil {
				return nil, err
			}

			res += a * b
			i += mulMatch[1]
		} else {
			i++
		}
	}

	return res, nil
}

func init() {
	registry.RegisterChallenge(3, Runner{})
}
