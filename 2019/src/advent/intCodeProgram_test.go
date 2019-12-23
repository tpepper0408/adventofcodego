package main

import (
	"fmt"
	"testing"
)

func Test_runProgram(t *testing.T) {
	input := []int{3, 0, 4, 0, 99}
	program := intCodeProgram(input)
	output, _, _ := program.runProgramOptionalDebug([]int{1}, true)
	if output != 1 {
		t.Error("Input-output - wrong output. Expected:", 1, " but got:", output)
	}
}
func Test_runProgramEqualTo8Position(t *testing.T) {
	output, _, _ := intCodeProgram{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}.runProgramOptionalDebug([]int{8}, true)
	if output != 1 {
		t.Error("Equal to 8 position - wrong output. Expected:", 1, " but got:", output)
	}
	output, _, _ = intCodeProgram{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}.runProgramOptionalDebug([]int{9}, true)
	if output != 0 {
		t.Error("Equal to 8 position - wrong output. Expected:", 0, " but got:", output)
	}
}
func Test_runProgramEqualTo8Value(t *testing.T) {
	output, _, _ := intCodeProgram{3, 3, 1108, -1, 8, 3, 4, 3, 99}.runProgramOptionalDebug([]int{8}, true)
	if output != 1 {
		t.Error("Equal to 8 value - wrong output. Expected:", 1, " but got:", output)
	}
	output, _, _ = intCodeProgram{3, 3, 1108, -1, 8, 3, 4, 3, 99}.runProgramOptionalDebug([]int{9}, true)
	if output != 0 {
		t.Error("Equal to 8 value - wrong output. Expected:", 0, " but got:", output)
	}
}

func Test_runProgramLessThan8Position(t *testing.T) {
	output, terminated, _ := intCodeProgram{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}.runProgramOptionalDebug([]int{8}, true)
	if output != 0 {
		t.Error("Less than 8 position - wrong output. Expected:", 0, " but got:", output)
	}
	if terminated {
		t.Error("Program should not have terminated")
	}
	output, terminated, _ = intCodeProgram{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}.runProgramOptionalDebug([]int{1}, true)
	if output != 1 {
		t.Error("Less than 8 position - wrong output. Expected:", 1, " but got:", output)
	}
	if terminated {
		t.Error("Program should not have terminated")
	}
}

func Test_runProgramLessThan8Value(t *testing.T) {
	output, terminated, _ := intCodeProgram{3, 3, 1107, -1, 8, 3, 4, 3, 99}.runProgramOptionalDebug([]int{8}, true)
	if output != 0 {
		t.Error("Less than 8 value - wrong output. Expected:", 0, " but got:", output)
	}
	if terminated {
		t.Error("Program should not have terminated")
	}
	output, terminated, _ = intCodeProgram{3, 3, 1107, -1, 8, 3, 4, 3, 99}.runProgramOptionalDebug([]int{1}, true)
	if output != 1 {
		t.Error("Less than 8 value - wrong output. Expected:", 1, " but got:", output)
	}
	if terminated {
		t.Error("Program should not have terminated")
	}
}
func Test_runProgramJumping(t *testing.T) {
	fmt.Println("Jump tests")
	output, terminated, _ := intCodeProgram{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}.runProgramOptionalDebug([]int{0}, true)
	if output != 0 {
		t.Error("Jump if false input 0 - wrong output. Expected:", 0, " but got:", output)
	}
	if terminated {
		t.Error("Program should not have terminated")
	}
	output, terminated, _ = intCodeProgram{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}.runProgramOptionalDebug([]int{1}, true)
	if output != 1 {
		t.Error("Jump if false input 1 - wrong output. Expected:", 0, " but got:", output)
	}
	if terminated {
		t.Error("Program should not have terminated")
	}
}

func Test_runProgramWithPhaseSetting43210(t *testing.T) {
	phaseSettings := []int{4, 3, 2, 1, 0}
	output := 0
	for _, p := range phaseSettings {
		output, _, _ = intCodeProgram{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}.runProgramOptionalDebug([]int{p, output}, true)
	}
	if output != 43210 {
		t.Error("43210 test failure - wrong output. Expected:", 43210, " but got:", output)
	}
}

func Test_runProgramWithPhaseSetting01234(t *testing.T) {
	phaseSettings := []int{0, 1, 2, 3, 4}
	output := 0
	for _, p := range phaseSettings {
		output, _, _ = intCodeProgram{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}.runProgramOptionalDebug([]int{p, output}, true)
	}
	if output != 54321 {
		t.Error("01234 test failure - wrong output. Expected:", 54321, " but got:", output)
	}
}

func Test_runProgramWithPhaseSetting10432(t *testing.T) {
	phaseSettings := []int{1, 0, 4, 3, 2}
	output := 0
	for _, p := range phaseSettings {
		output, _, _ = intCodeProgram{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
			1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}.runProgramOptionalDebug([]int{p, output}, true)
	}
	if output != 65210 {
		t.Error("10432 test failure - wrong output. Expected:", 65210, " but got:", output)
	}
}

func Test_runProgramWithoutTermination(t *testing.T) {
	output, terminated, _ := intCodeProgram{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}.runProgramOptionalDebug([]int{9, 0}, true)
	if terminated {
		t.Error("Should not have terminated")
	}
	if output != 5 {
		t.Error("Wrong output for termination test. Expected ", 5, " but got ", output)
	}
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
