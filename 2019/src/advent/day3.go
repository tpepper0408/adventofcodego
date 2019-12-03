package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"./utils"
)

type coordinate struct {
	x    int
	y    int
	path int
}

func dayThree() {
	lines := utils.ReadFile("Day3/input")
	var paths [][]coordinate
	var fullCoordinateList []coordinate
	for i, line := range lines {
		coordinates := processLine(line, i)
		paths = append(paths, coordinates)
		fullCoordinateList = append(fullCoordinateList, coordinates...)
	}
	crossingSections := findCrossingSections(fullCoordinateList)

	dayThreePartOne(crossingSections)
	dayThreePartTwo(paths, crossingSections)
}

func dayThreePartOne(crossingSections []coordinate) {
	fmt.Println("Crossing sections:", crossingSections)
	distances := findManhattanDistances(crossingSections)
	fmt.Println("Distances:", distances)
	fmt.Println("Min:", utils.FindMin(distances))
}

func dayThreePartTwo(paths [][]coordinate, crossingSections []coordinate) {
	var numberOfSteps []int
	for _, crossingCoordinate := range crossingSections {
		numberOfStepsForCrossing := 0
		for _, path := range paths {
			for i, step := range path {
				if step.x == crossingCoordinate.x &&
					step.y == crossingCoordinate.y {
					numberOfStepsForCrossing = numberOfStepsForCrossing + i
					continue
				}
			}
		}
		numberOfSteps = append(numberOfSteps, numberOfStepsForCrossing)
	}

	minSteps := utils.FindMin(numberOfSteps)
	fmt.Println("Minimum steps:", minSteps)
}

func processLine(line string, path int) []coordinate {
	directions := strings.Split(line, ",")
	//the first place is the centre
	coordinates := []coordinate{coordinate{x: 0, y: 0, path: path}}
	for _, direction := range directions {
		if strings.HasPrefix(direction, "R") {
			numberOfPlaces := strings.ReplaceAll(direction, "R", "")
			numberOfSteps, err := strconv.ParseInt(numberOfPlaces, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			lastCoordinate := coordinates[len(coordinates)-1]
			for index := 0; index < int(numberOfSteps); index++ {
				newCoordinate := coordinate{x: lastCoordinate.x + 1, y: lastCoordinate.y, path: lastCoordinate.path}
				coordinates = append(coordinates, newCoordinate)
				lastCoordinate = newCoordinate
			}
		} else if strings.HasPrefix(direction, "U") {
			numberOfPlaces := strings.ReplaceAll(direction, "U", "")
			numberOfSteps, err := strconv.ParseInt(numberOfPlaces, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			lastCoordinate := coordinates[len(coordinates)-1]
			for index := 0; index < int(numberOfSteps); index++ {
				newCoordinate := coordinate{x: lastCoordinate.x, y: lastCoordinate.y + 1, path: lastCoordinate.path}
				coordinates = append(coordinates, newCoordinate)
				lastCoordinate = newCoordinate
			}
		} else if strings.HasPrefix(direction, "L") {
			numberOfPlaces := strings.ReplaceAll(direction, "L", "")
			numberOfSteps, err := strconv.ParseInt(numberOfPlaces, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			lastCoordinate := coordinates[len(coordinates)-1]
			for index := 0; index < int(numberOfSteps); index++ {
				newCoordinate := coordinate{x: lastCoordinate.x - 1, y: lastCoordinate.y, path: lastCoordinate.path}
				coordinates = append(coordinates, newCoordinate)
				lastCoordinate = newCoordinate
			}
		} else if strings.HasPrefix(direction, "D") {
			numberOfPlaces := strings.ReplaceAll(direction, "D", "")
			numberOfSteps, err := strconv.ParseInt(numberOfPlaces, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			lastCoordinate := coordinates[len(coordinates)-1]
			for index := 0; index < int(numberOfSteps); index++ {
				newCoordinate := coordinate{x: lastCoordinate.x, y: lastCoordinate.y - 1, path: lastCoordinate.path}
				coordinates = append(coordinates, newCoordinate)
				lastCoordinate = newCoordinate
			}
		}
	}
	return coordinates
}

func findManhattanDistances(coordinates []coordinate) []int {
	manhattanDistances := make([]int, len(coordinates))
	for i, placement := range coordinates {
		manhattanDistance := utils.Abs(placement.x) + utils.Abs(placement.y)
		manhattanDistances[i] = manhattanDistance
	}
	return manhattanDistances
}

func findCrossingSections(coordinates []coordinate) []coordinate {
	var crossingSections []coordinate
	for _, pathCoordinate := range coordinates {
		//don't care about the centre
		if pathCoordinate.x == 0 &&
			pathCoordinate.y == 0 {
			continue
		}

		for _, pathCoordinate2 := range coordinates {
			//if on the same path then ignore
			if pathCoordinate.path == pathCoordinate2.path {
				continue
			}
			//don't care about the centre
			if pathCoordinate2.x == 0 &&
				pathCoordinate2.y == 0 {
				continue
			}
			if pathCoordinate.x == pathCoordinate2.x &&
				pathCoordinate.y == pathCoordinate2.y {
				crossingSections = append(crossingSections, pathCoordinate)
			}
		}
	}
	return crossingSections
}
