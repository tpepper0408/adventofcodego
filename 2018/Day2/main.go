package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// --- Day 2: Inventory Management System ---
	lines := readFile("input")
	partOne(lines)
}

func partOne(lines []string) {
	var numberOfTwoCharacterParcels int
	var numberOfThreeCharacterParcels int
	for _, line := range lines {
		var mapCharacters = make(map[rune]int)
		for _, c := range line {
			numCharacters := mapCharacters[c]
			numCharacters++
			mapCharacters[c] = numCharacters
		}

		var foundTwoCharacters bool
		var foundThreeCharacters bool
		for _, number := range mapCharacters {
			if number == 2 {
				foundTwoCharacters = true
			}
			if number == 3 {
				foundThreeCharacters = true
			}
		}

		if foundThreeCharacters {
			numberOfThreeCharacterParcels++
		}
		if foundTwoCharacters {
			numberOfTwoCharacterParcels++
		}
	}
	fmt.Println("Part one:")
	fmt.Println("Number of parcels with 2:", numberOfTwoCharacterParcels)
	fmt.Println("Number of parcels with 3:", numberOfThreeCharacterParcels)
	fmt.Println("Checksum:", numberOfTwoCharacterParcels*numberOfThreeCharacterParcels)
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
