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

	for _, instruction := range instructions {
		if instruction.Direction == "forward" {
			horiz += instruction.Distance
		} else if (instruction.Direction == "down") {
			vert += instruction.Distance
		} else if (instruction.Direction == "up") {
			vert -= instruction.Distance
		}
	}

	fmt.Printf("Horizontal: %v, Vertical %v, product of two: %v\n", horiz, vert, (horiz * vert))
}