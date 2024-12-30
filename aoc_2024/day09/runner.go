package day09

import (
	"bytes"
	"slices"
	"fmt"
	"log"
	"github.com/michelprogram/adventofcode/registry"
	"github.com/michelprogram/adventofcode/utils"
	"strconv"
)

type Runner struct{}

var _ utils.Challenge = (*Runner)(nil)

type Node struct {
	Id       int
	Value    int
	Type     bool
	Position int
}

type Disk struct {
	Data  []*Node
	File  []*Node
}

func NewDisk(data []byte) (*Disk, error) {

	inputs := bytes.TrimSpace(data)

	disk := &Disk{
		Data:  make([]*Node, 0),
		File:  make([]*Node, 0),
	}

	id := 0

	for i := 0; i < len(inputs); i++ {
		num, err := strconv.Atoi(string(inputs[i]))
		if err != nil {
			return nil, err
		}

		tmp := &Node{
			Value:    num,
			Position: i,
		}

		if i%2 == 0 {
			tmp.Id = id
			tmp.Type = true
			id++
			disk.File = append(disk.File, tmp)
		} else {
			tmp.Type = false
		}

		disk.Data = append(disk.Data, tmp)
	}

	return disk, nil
}

func (disk Disk) CheckSum() int{

	sum := 0
    position := 0
    
    for _, node := range disk.Data {
        if node.Type {
            for i := 0; i < node.Value; i++ {
                sum += position * node.Id
                position++
            }
        } else {
            position += node.Value
        }
    }
    
    return sum
}

func (disk Disk) String() (res string){

	for _, data := range disk.Data{
		for i := 0; i < data.Value; i++{
			if data.Type{
				res = fmt.Sprintf("%s%d", res,data.Id)
			}else{
				res = fmt.Sprintf("%s.", res)
			}
		}
	}

	return res

}

func (d Runner) ParseInput(data []byte) ([]int, error) {
	inputs := bytes.TrimSpace(data)
	res := make([]int, 0)

	id := 0

	for i := 0; i < len(inputs); i++ {
		num, err := strconv.Atoi(string(inputs[i]))
		if err != nil {
			return nil, err
		}

		if i%2 == 0 {
			for j := 0; j < num; j++ {
				res = append(res, id)
			}
			id++
		} else {
			for j := 0; j < num; j++ {
				res = append(res, -1)
			}
		}
	}

	return res, nil
}

func (d Runner) Part1(data []byte) (any, error) {
	disk, err := d.ParseInput(data)
	res := 0

	if err != nil {
		return nil, err
	}

	start := 0
	end := len(disk) - 1

	for start <= end {
		if disk[start] != -1 {
			res += start * disk[start]
			start++
			continue
		}
		if disk[end] == -1 {
			end--
			continue
		}

		res += start * disk[end]
		end--
		start++
	}

	return res, nil

}

func (d Runner) Part2(data []byte) (any, error) {

	disk, err := NewDisk(data)

	if err != nil {
		return nil, err
	}
	
	for i := len(disk.File) - 1; i > 0; i--{

		file := disk.File[i]
		position := file.Position

		for cursor := 0; cursor < position; cursor++{
			picked := disk.Data[cursor]
			
			if !picked.Type && picked.Value >= file.Value{
				disk.Data[cursor], disk.Data[position] = disk.Data[position], disk.Data[cursor]
				disk.Data[cursor].Position = cursor

				
				if diff := picked.Value - file.Value; diff > 0{
					disk.Data[position].Value = file.Value
					
					disk.Data = slices.Insert(disk.Data, cursor+1, &Node{
						Value:    diff,
						Position: cursor+1,
						Type: false,
					})

					for j := cursor+1; j < position ; j++{
						disk.Data[j].Position++
					}
				}
				
				break
			}
		}		
	}

	log.Println(disk)
	
	return disk.CheckSum(), nil
}

func init() {
	registry.RegisterChallenge(9, Runner{})
}
