package utils

import (
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FindMin(n []int) int {
	min := n[0]
	for _, value := range n {
		if value < min {
			min = value
		}
	}
	return min
}

func GetIntAsArrayOfInts(input int) []int {
	asString := strconv.Itoa(input)
	arrayOfLetters := strings.Split(asString, "")
	intsArray := []int{}
	for _, letter := range arrayOfLetters {
		i, _ := strconv.Atoi(letter)
		intsArray = append(intsArray, i)
	}
	return intsArray
}
