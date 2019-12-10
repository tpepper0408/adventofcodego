package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"./utils"
)

type intCodeProgram []int

const operationAdd int = 1
const operationMultiply int = 2
const operationInput int = 3
const operationOutput int = 4
const operationJumpIfTrue int = 5
const operationJumpIfFalse int = 6
const operationLessThan int = 7
const operationEquals int = 8

const param1 int = 2
const param2 int = 3
const param3 int = 4

func (program intCodeProgram) runProgram(initialInput []int) (int, bool) {
	return program.runProgramOptionalDebug(initialInput, false)
}
func (program intCodeProgram) runProgramOptionalDebug(initialInput []int, debug bool) (int, bool) {
	nextIndex := 0
	inputIndex := 0
	output := -1
	terminated := false
	for index := 0; true; index = nextIndex {
		opCode := program[index]
		if debug {
			fmt.Println("Opcode:", opCode)
		}
		if opCode < 0 {
			panic("Opcode should never be negative")
		}

		if isOperation(opCode, operationAdd) {
			p1 := program[index+1]
			if isPositionMode(opCode, param1) {
				p1 = program[program[index+1]]
			}
			p2 := program[index+2]
			if isPositionMode(opCode, param2) {
				p2 = program[program[index+2]]
			}
			valueToInsert := p1 + p2
			nextIndex = nextIndex + 4
			program[program[index+3]] = valueToInsert
		} else if isOperation(opCode, operationMultiply) {
			p1 := program[index+1]
			if isPositionMode(opCode, param1) {
				p1 = program[program[index+1]]
			}
			p2 := program[index+2]
			if isPositionMode(opCode, param2) {
				p2 = program[program[index+2]]
			}
			valueToInsert := p1 * p2
			nextIndex = nextIndex + 4
			program[program[index+3]] = valueToInsert
		} else if isOperation(opCode, operationInput) {
			if debug {
				fmt.Println("Got op 3 (value input)")
			}
			whereToStore := program[index+1]
			program[whereToStore] = initialInput[inputIndex]
			fmt.Println("Input:", initialInput[inputIndex], " to position ", whereToStore)
			inputIndex++
			nextIndex = nextIndex + 2
		} else if isOperation(opCode, operationOutput) {
			if debug {
				fmt.Println("Got op 4 (value output)")
			}
			out := program[index+1]
			if isPositionMode(opCode, param1) {
				out = program[program[index+1]]
			}
			inputIndex = 0
			output = out
			fmt.Println("Output:", out)
			nextIndex = nextIndex + 2
		} else if isOperation(opCode, operationJumpIfTrue) {
			toCheck := program[index+1]
			if isPositionMode(opCode, param1) {
				toCheck = program[program[index+1]]
			}
			if toCheck != 0 {
				valueToInsert := program[index+2]
				if isPositionMode(opCode, param2) {
					valueToInsert = program[program[index+2]]
				}
				program[index] = valueToInsert
				nextIndex = valueToInsert
			} else {
				nextIndex = nextIndex + 3
			}
		} else if isOperation(opCode, operationJumpIfFalse) {
			toCheck := program[index+1]
			if isPositionMode(opCode, param1) {
				toCheck = program[program[index+1]]
			}
			if toCheck == 0 {
				valueToInsert := program[index+2]
				if isPositionMode(opCode, param2) {
					valueToInsert = program[program[index+2]]
				}
				program[index] = valueToInsert
				nextIndex = valueToInsert
			} else {
				nextIndex = nextIndex + 3
			}
		} else if isOperation(opCode, operationLessThan) {
			toCheck1 := program[index+1]
			if isPositionMode(opCode, param1) {
				toCheck1 = program[program[index+1]]
			}
			toCheck2 := program[index+2]
			if isPositionMode(opCode, param2) {
				toCheck2 = program[program[index+2]]
			}
			valueToInsert := 0
			if toCheck1 < toCheck2 {
				valueToInsert = 1
			}
			program[program[index+3]] = valueToInsert
			nextIndex = nextIndex + 4
		} else if isOperation(opCode, operationEquals) {
			toCheck1 := program[index+1]
			if isPositionMode(opCode, param1) {
				toCheck1 = program[program[index+1]]
			}
			toCheck2 := program[index+2]
			if isPositionMode(opCode, param2) {
				toCheck2 = program[program[index+2]]
			}
			valueToInsert := 0
			if toCheck1 == toCheck2 {
				valueToInsert = 1
			}
			program[program[index+3]] = valueToInsert
			nextIndex = nextIndex + 4
		} else if opCode == 99 {
			if debug {
				fmt.Println("Terminated")
			}
			terminated = true
			break
		} else {
			panic("BOOM")
		}
	}

	if output == -1 {
		output = program[0]
	}
	return output, terminated
}

func isPositionMode(toCheck int, paramIndex int) bool {
	o := utils.GetIntAsArrayOfInts(toCheck)
	//if there isn't any mode information it is in position mode
	if len(o) <= 2 {
		return true
	}
	//if the index isn't there then it is in position mode
	if len(o)-1 < paramIndex {
		return true
	}
	//otherwise it should be 0
	return o[len(o)-1-paramIndex] == 0
}

func isOperation(toCheck int, op int) bool {
	if toCheck == op {
		return true
	}
	o := utils.GetIntAsArrayOfInts(toCheck)
	return o[len(o)-1] == op
}

func readInstructionLine(line string) []int {
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
