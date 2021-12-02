package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Direction string
	Distance int
}

func main()  {
	dat, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(dat), "\n")
	var instructions []Instruction

	for _, line := range lines {
		data := strings.Split(line, " ")
		dist, _ := strconv.Atoi(data[1])
		instructions = append(instructions, Instruction{
			Direction: data[0],
			Distance:  dist,
		})
	}

	horiz := 0
	vert := 0
	aim := 0

	for _, instruction := range instructions {
		if instruction.Direction == "forward" {
			horiz += instruction.Distance
			vert += instruction.Distance * aim
		} else if instruction.Direction == "down" {
			aim += instruction.Distance
		} else if instruction.Direction == "up" {
			aim -= instruction.Distance
		}
	}

	fmt.Printf("Horizontal: %v, Vertical %v, product of two: %v\n", horiz, vert, (horiz * vert))
}