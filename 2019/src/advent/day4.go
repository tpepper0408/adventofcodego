package main

import (
	"fmt"

	"./utils"
)

func dayFour() {
	fmt.Println("--- Day 4: Secure Container ---")
	minRange, maxRange := 256310, 732736
	dayFourPartOne(minRange, maxRange)
	dayFourPartTwo(minRange, maxRange)
}

func dayFourPartOne(minRange int, maxRange int) {
	var numbersThatFitPattern []int
	for index := minRange; index < maxRange; index++ {
		intsArray := utils.GetIntAsArrayOfInts(index)
		if !hasRepeatingInteger(intsArray) {
			continue
		}
		if !numbersAreSequential(intsArray) {
			continue
		}
		numbersThatFitPattern = append(numbersThatFitPattern, index)
	}
	fmt.Println("Part One - Number of potential passwords:", len(numbersThatFitPattern))
}

func dayFourPartTwo(minRange int, maxRange int) {
	var numbersThatFitPattern []int
	for index := minRange; index < maxRange; index++ {
		intsArray := utils.GetIntAsArrayOfInts(index)
		if !hasExactlyTwoRepeatingInteger(intsArray) {
			continue
		}
		if !numbersAreSequential(intsArray) {
			continue
		}
		numbersThatFitPattern = append(numbersThatFitPattern, index)
	}
	fmt.Println("Part Two - Number of potential passwords:", len(numbersThatFitPattern))
}

func numbersAreSequential(input []int) bool {
	lastInt := 0
	for _, number := range input {
		if number >= lastInt {
			lastInt = number
		} else {
			return false
		}
	}

	return true
}

func hasRepeatingInteger(input []int) bool {
	intsWithCounts := countInstancesOfRepeatingInts(input)
	for _, instances := range intsWithCounts {
		if instances > 1 {
			return true
		}
	}
	return false
}

func hasExactlyTwoRepeatingInteger(input []int) bool {
	intsWithCounts := countInstancesOfRepeatingInts(input)
	for _, instances := range intsWithCounts {
		if instances == 2 {
			return true
		}
	}
	return false
}

func countInstancesOfRepeatingInts(input []int) map[int]int {
	ints := make(map[int]int)

	for _, n := range input {
		numberOfInstances := ints[n]
		numberOfInstances = numberOfInstances + 1
		ints[n] = numberOfInstances
	}
	return ints
}
