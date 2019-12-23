package main

import "./utils"

func runRelay(program intCodeProgram, phaseSettings []int, debug bool) int {
	program1 := intCodeProgram(utils.DeepCopyInt(program))
	program2 := intCodeProgram(utils.DeepCopyInt(program))
	program3 := intCodeProgram(utils.DeepCopyInt(program))
	program4 := intCodeProgram(utils.DeepCopyInt(program))
	program5 := intCodeProgram(utils.DeepCopyInt(program))

	runningTotal := 0
	output, program1Terminated, program1Index := program1.runProgramOptionalDebug([]int{phaseSettings[0], runningTotal}, debug)
	if output != -1 {
		runningTotal = output
	}
	output, program2Terminated, program2Index := program2.runProgramOptionalDebug([]int{phaseSettings[1], runningTotal}, debug)
	if output != -1 {
		runningTotal = output
	}
	output, program3Terminated, program3Index := program3.runProgramOptionalDebug([]int{phaseSettings[2], runningTotal}, debug)
	if output != -1 {
		runningTotal = output
	}
	output, program4Terminated, program4Index := program4.runProgramOptionalDebug([]int{phaseSettings[3], runningTotal}, debug)
	if output != -1 {
		runningTotal = output
	}
	output, program5Terminated, program5Index := program5.runProgramOptionalDebug([]int{phaseSettings[4], runningTotal}, debug)
	if output != -1 {
		runningTotal = output
	}

	for {
		output, program1Terminated, program1Index = program1.runProgramOptionalDebugAndStartingIndex([]int{runningTotal}, debug, program1Index)
		if output != -1 {
			runningTotal = output
		}
		output, program2Terminated, program2Index = program2.runProgramOptionalDebugAndStartingIndex([]int{runningTotal}, debug, program2Index)
		if output != -1 {
			runningTotal = output
		}
		output, program3Terminated, program3Index = program3.runProgramOptionalDebugAndStartingIndex([]int{runningTotal}, debug, program3Index)
		if output != -1 {
			runningTotal = output
		}
		output, program4Terminated, program4Index = program4.runProgramOptionalDebugAndStartingIndex([]int{runningTotal}, debug, program4Index)
		if output != -1 {
			runningTotal = output
		}
		output, program5Terminated, program5Index = program5.runProgramOptionalDebugAndStartingIndex([]int{runningTotal}, debug, program5Index)
		if output != -1 {
			runningTotal = output
		}
		if program1Terminated &&
			program2Terminated &&
			program3Terminated &&
			program4Terminated &&
			program5Terminated {
			break
		}
	}
	return runningTotal
}
