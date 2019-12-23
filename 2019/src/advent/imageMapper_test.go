package main

import (
	"testing"

	"./utils"
)

func Test_mapImage(t *testing.T) {
	imageInts := utils.GetIntAsArrayOfInts(123456789012)
	result := mapImage(imageInts, 3, 2)
	result.print()
}

func Test_calculateChecksum(t *testing.T) {
	// layer := []int{1, 2, 0, 2, 1, 2, 2, 0, 1, 2, 2, 1, 2, 0, 2, 0, 2, 0, 1, 2, 1, 2, 1, 0, 0, 2, 0, 0, 2, 0, 2, 1, 0, 2, 0, 2, 2, 0, 1, 2, 2, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 2, 0, 1, 1, 0, 0, 1, 2, 2, 0, 1, 1, 1, 1, 0, 2, 0, 1, 2, 0, 1, 0, 2, 0, 2, 2, 2, 0, 2, 1, 2, 0, 2, 0, 2, 1, 0, 2, 0, 2, 2, 2, 2, 0, 2, 0, 1, 2, 2, 0, 1, 2, 0, 0, 0, 2, 0, 0, 2, 2, 0, 0, 0, 2, 2, 0, 0, 1, 0, 0, 2, 0, 1, 1, 1, 2, 2, 2, 0, 0, 0, 0, 1, 0, 0, 2, 0, 1, 2, 2, 1, 2, 1, 2, 1, 1}

	layer := []int{1, 1, 1, 2, 0, 0, 1, 1, 1, 2, 2, 2}
	result := calculateCheckSum(layer)
	if result != 4 {
		t.Error("Should be ", 4, " but is ", result)
	}
}
