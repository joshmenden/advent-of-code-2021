package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	dat, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(dat), "\n")
	fmt.Println(lines)
}