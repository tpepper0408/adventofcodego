package main

import (
	"fmt"
	"os"
	"strconv"

	"./utils"
)

func day1() {
	fmt.Println("--- Day 1: The Tyranny of the Rocket Equation ---")
	lines := utils.ReadFile("Day1/input")
	day1PartOne(lines)
	day1PartTwo(lines)
}

func day1PartOne(lines []string) {
	fmt.Println("Part One...")
	var totalValue int
	for _, line := range lines {
		intValue, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		intValue = (intValue / 3)
		intValue = intValue - 2
		fmt.Println("Value:", intValue)
		totalValue = totalValue + int(intValue)
	}
	fmt.Println("Total:", totalValue)
}

func day1PartTwo(lines []string) {
	fmt.Println("--- Part Two ---")
	var totalValue int
	for _, line := range lines {
		intValue, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fuelValue := calculateFuelValue(int(intValue))
		fmt.Println("Calculated Fuel:", fuelValue)
		totalValue = totalValue + int(fuelValue)
	}
	fmt.Println("Total:", totalValue)
}

func calculateFuelValue(massValue int) int {
	massValue = (massValue / 3)
	massValue = massValue - 2
	if massValue > 0 {
		extraMassValue := calculateFuelValue(massValue)
		if extraMassValue > 0 {
			massValue = massValue + extraMassValue
		}
	}
	return massValue
}
