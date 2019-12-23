package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(fileName string) []string {
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

func ReadStringLineToIntArray(line string) []int {
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
