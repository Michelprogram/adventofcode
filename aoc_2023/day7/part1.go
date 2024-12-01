package day7

import (
	"sort"
	"strconv"
	"strings"
)

var RANKS = map[uint8]int{
	50: 0,
	51: 1,
	52: 2,
	53: 3,
	54: 4,
	55: 5,
	56: 6,
	57: 7,
	84: 8,
	74: 9,
	81: 10,
	75: 11,
	65: 12,
}

type Hand struct {
	value string
	bid   int
}

func (g Hand) rank() (int, int) {

	var letters = make(map[int32]int)
	var max int

	for _, letter := range g.value {
		letters[letter] += 1
		if letters[letter] > max {
			max = letters[letter]
		}
	}

	return len(letters), max

}

type ByRank []Hand

func (hc ByRank) Len() int {
	return len(hc)
}

func (hc ByRank) Less(i, j int) bool {
	lenI, maxI := hc[i].rank()
	lenJ, maxJ := hc[j].rank()

	if lenI != lenJ {
		return lenI > lenJ
	}

	if maxI != maxJ {
		return maxI < maxJ
	}

	for k := 0; k < 5; k++ {
		rank1 := RANKS[hc[i].value[k]]
		rank2 := RANKS[hc[j].value[k]]
		if rank1 != rank2 {
			return rank1 < rank2
		}
	}

	return true

}

func (hc ByRank) Swap(i, j int) {
	hc[i], hc[j] = hc[j], hc[i]
}

func findHands(lines []string) (res []Hand) {

	for _, line := range lines {
		data := strings.Split(line, " ")

		r, _ := strconv.Atoi(data[1])

		res = append(res, Hand{
			value: data[0],
			bid:   r,
		})
	}

	return res

}

func cardsRank(data string) (res int) {

	lines := strings.Split(data, "\n")

	hands := findHands(lines)

	sort.Sort(ByRank(hands))

	res = hands[0].bid

	for i := 1; i < len(hands); i++ {
		res += hands[i].bid * (i + 1)
	}

	return res
}
