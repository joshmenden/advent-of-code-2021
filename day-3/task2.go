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

	oxygenRating := strings.Join(filterData(lines, 0, "most"), "")
	co2Rating := strings.Join(filterData(lines, 0, "least"), "")
	fmt.Println(oxygenRating)
	fmt.Println(co2Rating)

	oxygenDec, _ := strconv.ParseInt(oxygenRating, 2, 64)
	co2Dec, _ := strconv.ParseInt(co2Rating, 2, 64)
	fmt.Println(oxygenDec)
	fmt.Println(co2Dec)

	fmt.Println(oxygenDec * co2Dec)
}

func filterData(dataSet []string, digitIndex int, comparator string) []string {
	if len(dataSet) == 1 {
		return dataSet
	}

	zeroes := []string{}
	ones := []string{}
	for j := 0; j < len(dataSet); j++ {
		bit := string(dataSet[j][digitIndex])
		if bit == "0" {
			zeroes = append(zeroes, dataSet[j])
		} else if bit == "1" {
			ones = append(ones, dataSet[j])
		}
	}

	if comparator == "most" {
		if len(zeroes) == len(ones) {
			return filterData(ones, digitIndex + 1, comparator)
		} else if len(zeroes) > len(ones) {
			return filterData(zeroes, digitIndex + 1, comparator)
		} else {
			return filterData(ones, digitIndex + 1, comparator)
		}
	} else { // most be least
		if len(zeroes) == len(ones) {
			return filterData(zeroes, digitIndex+1, comparator)
		} else if len(zeroes) < len(ones) {
			return filterData(zeroes, digitIndex + 1, comparator)
		} else {
			return filterData(ones, digitIndex + 1, comparator)
		}
	}

}