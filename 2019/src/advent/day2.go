package main

import (
	"fmt"

	"./utils"
)

func day2() {
	lines := utils.ReadFile("Day2/input")
	instructions := readInstructionLine(lines[0])
	day2PartOne(instructions)
	day2PartTwo(instructions)
}

func day2PartOne(instructions []int) int {
	fmt.Println("Starting program:", instructions)

	tempInstructions := utils.DeepCopyInt(instructions)
	tempInstructions[1] = 12
	tempInstructions[2] = 2
	fmt.Println("Applied 1202 program alarm")

	finishedProgram, _ := intCodeProgram(tempInstructions).runProgram([]int{0})
	fmt.Println("Value in position 0: ", finishedProgram)
	return finishedProgram
}

func day2PartTwo(instructions []int) (int, int) {
	fmt.Println("--- Part Two ---")
	fmt.Println("Looking for: ", 19690720)

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			if runWithReplacedValues(instructions, i, j) {
				return i, j
			}

		}
	}
	return -1, -1
}

func runWithReplacedValues(program []int, noun int, verb int) bool {
	tempInstructions := utils.DeepCopyInt(program)
	tempInstructions[1] = noun
	tempInstructions[2] = verb
	finishedProgram, _ := intCodeProgram(tempInstructions).runProgram([]int{0})
	if finishedProgram == 19690720 {
		fmt.Println("Found values: noun - ", noun, " verb - ", verb)
		return true
	}
	return false
}
