package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"./utils"
)

func day2() {
	lines := utils.ReadFile("Day2/input")
	instructions := readLine(lines[0])
	day2PartOne(instructions)
	day2PartTwo(instructions)
}

func day2PartOne(instructions []int) {
	fmt.Println("Starting program:", instructions)

	tempInstructions := utils.DeepCopyInt(instructions)
	tempInstructions[1] = 12
	tempInstructions[2] = 2
	fmt.Println("Applied 1202 program alarm")

	finishedProgram := runProgram(tempInstructions)
	fmt.Println("Value in position 0: ", finishedProgram[0])
}

func day2PartTwo(instructions []int) {
	fmt.Println("--- Part Two ---")
	fmt.Println("Looking for: ", 19690720)

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			runWithReplacedValues(instructions, i, j)
		}
	}
}

func runWithReplacedValues(program []int, noun int, verb int) {
	tempInstructions := utils.DeepCopyInt(program)
	tempInstructions[1] = noun
	tempInstructions[2] = verb
	finishedProgram := runProgram(tempInstructions)
	if finishedProgram[0] == 19690720 {
		fmt.Println("Found values: noun - ", noun, " verb - ", verb)
		os.Exit(0)
	}
}

func runProgram(program []int) []int {
	for index := 0; true; index = index + 4 {
		operation := program[index]
		var valueToInsert int
		if operation == 1 {
			valueToInsert = program[program[index+1]] + program[program[index+2]]
		} else if operation == 2 {
			valueToInsert = program[program[index+1]] * program[program[index+2]]
		} else if operation == 99 {
			break
		}
		program[program[index+3]] = valueToInsert
	}
	return program
}

func readLine(line string) []int {
	instructions := strings.Split(line, ",")
	var intInstructions []int
	for _, instruction := range instructions {
		intValue, err := strconv.ParseInt(instruction, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		intInstructions = append(intInstructions, int(intValue))
	}
	return intInstructions
}
