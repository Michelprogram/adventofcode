package aoc2024

import (
	"bytes"
	"github.com/michelprogram/adventofcode/utils"
	"strconv"
)

type Day9 struct{}

var _ utils.Challenge = (*Day9)(nil)

func (d Day9) ParseInputs(data []byte) ([]int, error) {
	inputs := bytes.TrimSpace(data)
	var result []int

	for i := 0; i < len(inputs); i++ {
		num, err := strconv.Atoi(string(inputs[i]))
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}

	return result, nil
}

func sum(numbers []int) (res int) {
	for _, number := range numbers {
		res += number
	}
	return res
}

func fillArrayWithId(inputs []int) []int {
	id := 0
	containerCursor := 0
	blocks := make([]int, sum(inputs))

	for i, count := range inputs {
		for j := 0; j < count; j++ {
			if i%2 == 0 {
				blocks[containerCursor] = id
			} else {
				blocks[containerCursor] = -1
			}
			containerCursor++
		}
		if i%2 == 0 {
			id++
		}
	}

	return blocks
}

type Disk struct {
	Map    []int
	Blocks []int
}

func NewDisk(inputs []int) *Disk {
	return &Disk{
		Map:    inputs,
		Blocks: fillArrayWithId(inputs),
	}
}

func (d Disk) indexFreeSpace() (start int) {
	for i, block := range d.Blocks {
		if block == -1 {
			return i
		}
	}
	return len(d.Blocks)
}

func (d *Disk) Compress() {

	start := d.indexFreeSpace()
	end := len(d.Blocks) - 1

	for start < end {
		if d.Blocks[start] != -1 {
			start++
			continue
		}
		if d.Blocks[end] == -1 {
			end--
			continue
		}

		// Swap the blocks
		d.Blocks[start], d.Blocks[end] = d.Blocks[end], d.Blocks[start]
		start++
		end--
	}

	d.Blocks = d.Blocks[:d.indexFreeSpace()]

}

func (d *Disk) CompressV2() {

	start := d.indexFreeSpace()
	end := len(d.Blocks) - 1

	for start < end {
		if d.Blocks[start] != -1 {
			start++
			continue
		}
		if d.Blocks[end] == -1 {
			end--
			continue
		}

		// Swap the blocks
		d.Blocks[start], d.Blocks[end] = d.Blocks[end], d.Blocks[start]
		start++
		end--
	}

	d.Blocks = d.Blocks[:d.indexFreeSpace()]

}

func (d Disk) CheckSum() (res int) {

	for i, block := range d.Blocks {
		res += i * block
	}

	return res
}

func (d Day9) Part1(data []byte) (any, error) {

	inputs, err := d.ParseInputs(data)

	if err != nil {
		return nil, err
	}

	disk := NewDisk(inputs)

	disk.Compress()

	return disk.CheckSum(), nil
}

func (d Day9) Part2(data []byte) (any, error) {

	return nil, nil
}
