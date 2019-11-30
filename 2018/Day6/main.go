package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type coordinate struct {
	x    int
	y    int
	code string
}
type areaCode struct {
	code          string
	expansions    int
	distanceScore int
}
type grid [][]areaCode

var maxDistanceScore int
var usedCodes []string

func main() {
	lines := readFile("input")

	fmt.Println("--- Day 6: Chronal Coordinates ---")
	//figure out how big the grid should be and the initial coordinates
	largestX, largestY, initialCoords := parseInputLines(lines)

	//assign area codes to the areas
	initialCoords = assignAreaCodes(initialCoords)
	for _, coord := range initialCoords {
		coord.print()
	}

	//now build the grid with the initial coordinates
	builtGrid := buildGrid(largestX, largestY, initialCoords)
	fmt.Println("Initial state:")
	builtGrid.print()

	partOne(builtGrid)
	partTwo(builtGrid, initialCoords)
}

func partOne(g grid) {
	fmt.Println("Part one...")
	iteration := 1
	for {
		if checkFinished(g) {
			break
		}
		g = propagate(g, iteration)
		iteration++
	}
	g.print()

	codesToExclude := findInfiniteZones(g)
	fmt.Println("Codes to exclude: ", codesToExclude)

	//now find the areas that are not infinite
	areaMap := make(map[string]int)
	for _, column := range g {
		for _, areaCode := range column {
			code := strings.ToLower(areaCode.code)
			if !contains(codesToExclude, code) {
				number := areaMap[code]
				number = number + 1
				areaMap[code] = number
			}
		}
	}
	fmt.Println("Areas:", areaMap)
	biggestArea := 0
	for _, area := range areaMap {
		if area > biggestArea {
			biggestArea = area
		}
	}
	fmt.Println("Biggest area: ", biggestArea)
}

func partTwo(g grid, coords []coordinate) {
	fmt.Println("Part two...")
	maxDistanceScore = 10000

	for y, column := range g {
		for x, areaCode := range column {
			distanceScore := 0
			for _, coord := range coords {
				score := absolute(x-coord.x) + absolute(y-coord.y)
				distanceScore = distanceScore + score
			}
			areaCode.distanceScore = distanceScore
			g[y][x] = areaCode
		}
	}
	g.print()

	numberWithinArea := 0
	for _, column := range g {
		for _, areaCode := range column {
			if areaCode.distanceScore < maxDistanceScore {
				numberWithinArea++
			}
		}
	}
	fmt.Println("Number in area: ", numberWithinArea)
}

func absolute(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func findInfiniteZones(g grid) []string {
	infiniteCodes := []string{}
	for _, column := range g {
		codeToExcludeLeft := column[0]
		codeToExcludeRight := column[len(column)-1]
		infiniteCodes = append(infiniteCodes, codeToExcludeLeft.code, codeToExcludeRight.code)
	}
	firstLine := g[0]
	for _, areaCode := range firstLine {
		infiniteCodes = append(infiniteCodes, areaCode.code)
	}
	lastLine := g[len(g)-1]
	for _, areaCode := range lastLine {
		infiniteCodes = append(infiniteCodes, areaCode.code)
	}

	return removeDuplicates(infiniteCodes)
}

func removeDuplicates(list []string) []string {
	keys := make(map[string]bool)
	distinct := []string{}
	for _, entry := range list {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			distinct = append(distinct, entry)
		}
	}
	return distinct
}

func propagate(g grid, iteration int) grid {
	//first find the coordinates to work with
	coordinates := []coordinate{}
	for i, column := range g {
		for j, position := range column {
			if position.code != "" {
				coordinates = append(coordinates, coordinate{x: j, y: i, code: position.code})
			}
		}
	}

	//now propogate out in each direction
	for _, coord := range coordinates {
		//left
		if coord.x-1 >= 0 {
			areaCode := g[coord.y][coord.x-1]
			if areaCode.expansions == iteration &&
				areaCode.code != strings.ToLower(coord.code) {
				areaCode.code = "."
			} else if areaCode.code == "" {
				areaCode.code = strings.ToLower(coord.code)
				areaCode.expansions = iteration
			}
			g[coord.y][coord.x-1] = areaCode
		}
		//top
		if coord.y-1 >= 0 {
			areaCode := g[coord.y-1][coord.x]
			if areaCode.expansions == iteration &&
				areaCode.code != strings.ToLower(coord.code) {
				areaCode.code = "."
			} else if areaCode.code == "" {
				areaCode.code = strings.ToLower(coord.code)
				areaCode.expansions = iteration
			}
			g[coord.y-1][coord.x] = areaCode
		}
		//right
		if coord.x+1 < len(g[0]) {
			areaCode := g[coord.y][coord.x+1]
			if areaCode.expansions == iteration &&
				areaCode.code != strings.ToLower(coord.code) {
				areaCode.code = "."
			} else if areaCode.code == "" {
				areaCode.code = strings.ToLower(coord.code)
				areaCode.expansions = iteration
			}
			g[coord.y][coord.x+1] = areaCode
		}
		//bottom
		if coord.y+1 < len(g) {
			areaCode := g[coord.y+1][coord.x]
			if areaCode.expansions == iteration &&
				areaCode.code != strings.ToLower(coord.code) {
				areaCode.code = "."
			} else if areaCode.code == "" {
				areaCode.code = strings.ToLower(coord.code)
				areaCode.expansions = iteration
			}
			g[coord.y+1][coord.x] = areaCode
		}
	}
	return g
}

func checkFinished(g grid) bool {
	for _, column := range g {
		for _, areaCode := range column {
			if areaCode.code == "" {
				return false
			}
		}
	}
	return true
}

func assignAreaCodes(coords []coordinate) []coordinate {
	for i, coord := range coords {
		coord.code = generateCode()
		coords[i] = coord
	}
	return coords
}

func generateCode() string {
	possibleCodes := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
		"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V",
		"X", "Y", "Z", "AA", "BB", "CC", "DD", "EE", "FF", "GG",
		"HH", "II", "JJ", "KK", "LL", "MM", "NN", "OO", "PP", "QQ",
		"RR", "SS", "TT", "UU", "VV", "XX", "YY", "ZZ",
	}

	for _, codeToUse := range possibleCodes {
		if contains(usedCodes, codeToUse) {
			continue
		}
		usedCodes = append(usedCodes, codeToUse)
		return codeToUse
	}

	return ""
}

func contains(l []string, s string) bool {
	for _, toCheck := range l {
		if s == toCheck {
			return true
		}
	}
	return false
}

func buildGrid(maxX int, maxY int, coordinates []coordinate) grid {
	//we want to build a grid with an extra layer around the outside
	gridBuild := make(grid, maxY+2)
	for index := 0; index < len(gridBuild); index++ {
		gridBuild[index] = make([]areaCode, maxX+2)
	}

	for _, coord := range coordinates {
		gridBuild[coord.y][coord.x].code = coord.code
	}
	return gridBuild
}

func (g grid) print() {
	fmt.Println("x:", len(g[0]))
	fmt.Println("y:", len(g))
	//header
	fmt.Print(".\t")
	for index := 0; index < len(g[0]); index++ {
		fmt.Print(index)
		fmt.Print("  ")
	}
	fmt.Println()

	for i, column := range g {
		fmt.Print(i)
		fmt.Print("\t")
		for _, placement := range column {
			fmt.Print(placement.getAreaCode())
			fmt.Print("  ")
		}
		fmt.Println()
	}
}

func (c coordinate) print() {
	fmt.Printf("coordinate: %+v", c)
	fmt.Println()
}

func (a areaCode) getAreaCode() string {
	if a.distanceScore < maxDistanceScore &&
		!isOriginalAreaCode(a.code) {
		return "#"
	}

	if a.code == "" {
		return "_"
	}
	return a.code
}

func isOriginalAreaCode(s string) bool {
	r := []rune(s)
	return unicode.IsUpper(r[0])
}

func parseInputLines(lines []string) (int, int, []coordinate) {
	var largestX int
	var largestY int
	initialCoords := []coordinate{}
	for _, line := range lines {
		coords := strings.Split(line, ", ")
		x, err := strconv.ParseInt(coords[0], 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		y, err2 := strconv.ParseInt(coords[1], 10, 64)
		if err2 != nil {
			fmt.Println("Error:", err2)
			os.Exit(1)
		}
		initialCoords = append(initialCoords, coordinate{x: int(x), y: int(y)})
		if int(x) > largestX {
			largestX = int(x)
		}
		if int(y) > largestY {
			largestY = int(y)
		}
	}
	return largestX, largestY, initialCoords
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
