package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := os.Args[1]
	lines := readFile(fileName)
	m := make(map[int64]bool)
	var frequency int64
	frequency = 0
	for {
		for _, line := range lines {
			addition := strings.HasPrefix(line, "+")
			delta := findFrequency(line)
			if addition {
				frequency = frequency + delta
			} else {
				frequency = frequency - delta
			}

			if m[frequency] {
				fmt.Println("Found repeating frequency", frequency)
				os.Exit(0)
			} else {
				m[frequency] = true
			}
		}
	}
}

func findFrequency(line string) int64 {
	noAddSign := strings.Replace(line, "+", "", 1)
	noMinusSign := strings.Replace(noAddSign, "-", "", 1)
	value, err := strconv.ParseInt(noMinusSign, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return value
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
