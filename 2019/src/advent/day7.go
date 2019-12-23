package main

import (
	"fmt"

	"./utils"
)

func day7() {
	lines := utils.ReadFile("Day7/input")
	program := intCodeProgram(utils.ReadStringLineToIntArray(lines[0]))
	day7PartOne(program)
	day7PartTwo(program)
}

func day7PartOne(p intCodeProgram) {
	phaseSettings := utils.GetPermutations([]int{0, 1, 2, 3, 4})
	largestOutput := 0
	for _, phaseSetting := range phaseSettings {
		fmt.Println("Got phase setting:", phaseSetting)
		output := 0
		for _, s := range phaseSetting {
			programToRun := intCodeProgram(utils.DeepCopyInt(p))
			output, _, _ = programToRun.runProgramOptionalDebug([]int{s, output}, false)
		}
		if output > largestOutput {
			largestOutput = output
		}
	}
	fmt.Println("Largest output:", largestOutput)
}

func day7PartTwo(p intCodeProgram) {
	phaseSettings := utils.GetPermutations([]int{9, 8, 7, 6, 5})
	largestOutput := 0
	for _, phaseSetting := range phaseSettings {
		fmt.Println("Got phase setting:", phaseSetting)
		output := runRelay(p, phaseSetting, false)
		if output > largestOutput {
			largestOutput = output
		}
	}
	fmt.Println("Largest output:", largestOutput)
}
