package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	dat, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(dat), "\n")

	gammaBit := []string{}
	epsilonBit := []string{}

	for i := 0; i < len(lines[0]); i++ {
		zeroes := 0
		ones := 0
		for j := 0; j < len(lines); j++ {
			bit := string(lines[j][i])
			if bit == "0" {
				 zeroes += 1
			} else if bit == "1" {
				 ones += 1
			}
		}

		if zeroes > ones {
			gammaBit = append(gammaBit, "0")
			epsilonBit = append(epsilonBit, "1")
		} else {
			gammaBit = append(gammaBit, "1")
			epsilonBit = append(epsilonBit, "0")
		}
	}



	gammaBitStr := strings.Join(gammaBit, "")
	epsilonBitStr := strings.Join(epsilonBit, "")

	gammaDecimal, _ := strconv.ParseInt(gammaBitStr, 2, 64)
	epsilonDecimal, _ := strconv.ParseInt(epsilonBitStr, 2, 64)

	fmt.Println(gammaDecimal * epsilonDecimal)
}