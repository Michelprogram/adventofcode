package day8

import (
	"strings"
)

var (
	L_UINT8           uint8 = 76
	INDEX_NAVIGATE          = 3
	INDEX_LEFT_START        = 7
	INDEX_LEFT_END          = 10
	INDEX_RIGHT_START       = 12
	INDEX_RIGHT_END         = 15
	GOAL                    = "ZZZ"
)

type Node struct {
	value string
	right string
	left  string
}

func camelDirection(data string) int {

	var nodes = make(map[string]Node)
	var node Node
	var iteration int

	splitted := strings.Split(data, "\n\n")

	directions, camels := splitted[0], splitted[1]

	lines := strings.Split(camels, "\n")

	for i := len(lines) - 1; i >= 0; i-- {
		camel := lines[i]

		value := camel[:INDEX_NAVIGATE]

		node = Node{
			value: value,
			left:  camel[INDEX_LEFT_START:INDEX_LEFT_END],
			right: camel[INDEX_RIGHT_START:INDEX_RIGHT_END],
		}

		nodes[value] = node
	}

	size := len(directions)

	node = nodes["AAA"]

	for node.value != GOAL {

		direction := directions[iteration%size]

		if direction == L_UINT8 {
			node = nodes[node.left]
		} else {
			node = nodes[node.right]
		}

		iteration++
	}

	return iteration

}
