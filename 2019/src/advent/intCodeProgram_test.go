package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_runProgram(t *testing.T) {
	input := []int{3, 0, 4, 0, 99}
	program := intCodeProgram(input)
	runProgram := program.runProgramOptionalDebug(1, true)
	if !reflect.DeepEqual(runProgram, input) {
		t.Error("Wrong ouput. Expected:", input, " but got:", program)
	}

	fmt.Println("Equal to 8")
	input = []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	program = intCodeProgram(input)
	runProgram = program.runProgramOptionalDebug(8, true)
	runProgram = program.runProgramOptionalDebug(9, true)

	program = intCodeProgram{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	runProgram = program.runProgramOptionalDebug(8, true)
	runProgram = program.runProgramOptionalDebug(9, true)

	fmt.Println("Less than 8")
	program = intCodeProgram{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	runProgram = program.runProgramOptionalDebug(8, true)
	runProgram = program.runProgramOptionalDebug(1, true)
	intCodeProgram{3, 3, 1107, -1, 8, 3, 4, 3, 99}.runProgramOptionalDebug(8, true)
	intCodeProgram{3, 3, 1107, -1, 8, 3, 4, 3, 99}.runProgramOptionalDebug(1, true)

	fmt.Println("Jump tests")
	intCodeProgram{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}.runProgramOptionalDebug(0, true)
	intCodeProgram{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}.runProgramOptionalDebug(1, true)
}

func Test_isOperation(t *testing.T) {

	if !isOperation(4, operationOutput) {
		t.Error("Was expecting ", 4, " to be output operation.")
	}
	if !isOperation(1004, operationOutput) {
		t.Error("Was expecting ", 1004, " to be output operation.")
	}
}

func Test_isPositionMode(t *testing.T) {
	if isPositionMode(1101, param1) {
		t.Error("Expecting immediate mode for param1 in ", 1101)
	}
	if isPositionMode(1101, param2) {
		t.Error("Expecting immediate mode for param2 in ", 1101)
	}
	if !isPositionMode(1101, param3) {
		t.Error("Expecting position mode for param3 in ", 1101)
	}

	if isPositionMode(1001, param2) {
		t.Error("Expecting immediate mode for param2 in ", 1001)
	}
	if !isPositionMode(1001, param1) {
		t.Error("Expecting position mode for param1 in ", 1001)
	}

	if !isPositionMode(1002, param1) {
		t.Error("Expecting position mode for param1 in ", 1002)
	}
	if isPositionMode(1002, param2) {
		t.Error("Expecting immediate mode for param2 in ", 1002)
	}
}
