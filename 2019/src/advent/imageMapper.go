package main

import "fmt"

type imageMap [][]int

func mapImage(imageInts []int, pixelWidth int, pixelHeight int) imageMap {
	layerSize := pixelWidth * pixelHeight
	numberLayers := len(imageInts) / layerSize

	startIndex := 0
	endIndex := layerSize

	retval := make([][]int, numberLayers)
	for index := 0; index < numberLayers; index++ {
		retval[index] = imageInts[startIndex:endIndex]
		startIndex = startIndex + layerSize
		endIndex = endIndex + layerSize
	}
	return imageMap(retval)
}

func (i imageMap) print() {
	for i, layer := range i {
		fmt.Println("Layer", i, ":", layer)
	}
}

func (i imageMap) findLayerWithFewestZeroes() int {
	smallestNumberOfZeroes := len(i[0])
	indexOfSmallestNumber := -1
	for index, layer := range i {
		numberOfZeroes := 0
		for _, l := range layer {
			if l == 0 {
				numberOfZeroes++
			}
		}
		fmt.Println("Found ", numberOfZeroes, " zeroes for layer ", index)

		if numberOfZeroes < smallestNumberOfZeroes {
			smallestNumberOfZeroes = numberOfZeroes
			indexOfSmallestNumber = index
		}
	}
	fmt.Println("Found smallest number of zeroes ", smallestNumberOfZeroes, " in layer ", indexOfSmallestNumber)
	fmt.Println("Layer: ", i[indexOfSmallestNumber])
	return indexOfSmallestNumber
}

func (i imageMap) calculateChecksumForLayer(l int) int {
	layer := i[l]
	return calculateCheckSum(layer)
}

func calculateCheckSum(layer []int) int {
	numberOfOnes := 0
	numberOfTwos := 0
	for _, bit := range layer {
		if bit == 1 {
			numberOfOnes++
		} else if bit == 2 {
			numberOfTwos++
		}
	}
	fmt.Println("Number of ones:", numberOfOnes, " and twos:", numberOfTwos)
	return numberOfOnes * numberOfTwos
}
