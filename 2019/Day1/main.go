package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("--- Day 1: The Tyranny of the Rocket Equation ---")
	lines := readFile("input")
	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
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

func partTwo(lines []string) {
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

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
