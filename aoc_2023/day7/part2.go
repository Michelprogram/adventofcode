package day7

import (
	"sort"
	"strings"
)

var JOKER8 uint8 = 74
var JOKER32 int32 = 74

var RANKS_JOKER = map[uint8]int{
	JOKER8: 0,
	50:     1,
	51:     2,
	52:     3,
	53:     4,
	54:     5,
	55:     6,
	56:     7,
	57:     8,
	84:     9,
	81:     10,
	75:     11,
	65:     12,
}

type ByRankJoker []Hand

func (hc ByRankJoker) Len() int {
	return len(hc)
}

func (hc ByRankJoker) Less(i, j int) bool {
	lenI, maxI := hc[i].rankJoker()
	lenJ, maxJ := hc[j].rankJoker()

	if lenI != lenJ {
		return lenI > lenJ
	}

	if maxI != maxJ {
		return maxI < maxJ
	}

	for k := 0; k < 5; k++ {
		rank1 := RANKS_JOKER[hc[i].value[k]]
		rank2 := RANKS_JOKER[hc[j].value[k]]
		if rank1 != rank2 {
			return rank1 < rank2
		}
	}

	return true

}

func (hc ByRankJoker) Swap(i, j int) {
	hc[i], hc[j] = hc[j], hc[i]
}

func (g Hand) rankJoker() (int, int) {

	var letters = make(map[int32]int)
	var max int

	for _, letter := range g.value {
		letters[letter] += 1
		if letters[letter] > max && letter != JOKER32 {
			max = letters[letter]
		}
	}

	if letters[JOKER32] == 5 {
		return len(letters) + 1, letters[JOKER32]
	}

	if letters[JOKER32] > 0 {
		return len(letters) - 1, letters[JOKER32] + max
	}

	return len(letters), letters[JOKER32] + max

}

func cardsRankJoker(data string) (res int) {

	lines := strings.Split(data, "\n")

	hands := findHands(lines)

	sort.Sort(ByRankJoker(hands))
	res = hands[0].bid

	for i := 1; i < len(hands); i++ {
		res += hands[i].bid * (i + 1)
	}

	return res
}
