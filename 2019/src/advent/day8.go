package main

import (
	"fmt"
	"strconv"

	"./utils"
)

func day8() {
	input := utils.ReadFile("Day8/input")[0]

	numbers := make([]int, len(input))
	for i, char := range input {
		numbers[i], _ = strconv.Atoi(string(char))
	}

	day8PartOne(numbers)
}

func day8PartOne(input []int) {
	image := mapImage(input, 25, 6)
	image.print()
	layerIndex := image.findLayerWithFewestZeroes()
	fmt.Println("Layer with most zeroes:", layerIndex)
	checkSum := image.calculateChecksumForLayer(layerIndex)
	fmt.Println("Checksum:", checkSum)
}
