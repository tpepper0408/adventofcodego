package utils

import (
	"bufio"
	"fmt"
	"os"
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
