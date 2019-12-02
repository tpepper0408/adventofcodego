package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("input")
	instructions := readLine(lines[0])
	partOne(instructions)
	partTwo(instructions)
}

func partOne(instructions []int) {
	fmt.Println("Starting program:", instructions)

	tempInstructions := deepCopy(instructions)
	tempInstructions[1] = 12
	tempInstructions[2] = 2
	fmt.Println("Applied 1202 program alarm")

	finishedProgram := runProgram(tempInstructions)
	fmt.Println("Value in position 0: ", finishedProgram[0])
}

func partTwo(instructions []int) {
	fmt.Println("--- Part Two ---")
	fmt.Println("Looking for: ", 19690720)

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			runWithReplacedValues(instructions, i, j)
		}
	}
}

func runWithReplacedValues(program []int, noun int, verb int) {
	tempInstructions := deepCopy(program)
	tempInstructions[1] = noun
	tempInstructions[2] = verb
	finishedProgram := runProgram(tempInstructions)
	if finishedProgram[0] == 19690720 {
		fmt.Println("Found values: noun - ", noun, " verb - ", verb)
		os.Exit(0)
	}
}

func deepCopy(original []int) []int {
	copy := make([]int, len(original))
	for i, originalValue := range original {
		copy[i] = originalValue
	}
	return copy
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
