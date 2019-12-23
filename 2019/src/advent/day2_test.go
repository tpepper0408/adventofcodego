package main

import (
	"testing"

	"./utils"
)

func Test_day2PartOne(t *testing.T) {
	lines := utils.ReadFile("Day2/input")
	instructions := utils.ReadStringLineToIntArray(lines[0])
	result := day2PartOne(instructions)
	if result != 3224742 {
		t.Error("Wrong output. Expected ", 3224742, " but got ", result)
	}
}

func Test_day2PartTwo(t *testing.T) {
	lines := utils.ReadFile("Day2/input")
	instructions := utils.ReadStringLineToIntArray(lines[0])
	noun, verb := day2PartTwo(instructions)
	if noun != 79 {
		t.Error("Wrong noun. Expected ", 79, " but got ", noun)
	}
	if verb != 60 {
		t.Error("Wrong verb. Expected ", 60, " but got ", verb)
	}
}
