package main

import (
	"./utils"
)

func day5() {
	lines := utils.ReadFile("Day5/input")
	program := readInstructionLine(lines[0])
	programToRun := intCodeProgram(program)
	day5PartOne(programToRun)
	day5PartTwo(programToRun)
}

func day5PartOne(programToRun intCodeProgram) {
	programToRun.runProgram([]int{1})
}

func day5PartTwo(programToRun intCodeProgram) {
	programToRun.runProgram([]int{5})
}
