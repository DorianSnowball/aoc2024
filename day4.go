package main

import (
	"bufio"
	"os"
	"time"
)

func day4() {
	println("Solving AoC 2024 Day 4:")
	input := loadInputDay4()
	day4part1(input)
	day4part2(input)
}

func loadInputDay4() []string {
	defer timeTrack(time.Now(), "Loading/Parsing Input")
	file, err := os.Open("inputs/day4.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

var maxX int
var maxY int

func day4part1(lines []string) {
	defer timeTrack(time.Now(), "Part1")
	var totalFound = 0

	maxX = len(lines[0])
	maxY = len(lines)

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			char := lines[x][y]
			if char == 'X' {
				totalFound += checkXMASSpot(lines, x, y)
			}
		}
	}

	println(totalFound)
}

func checkXMASSpot(lines []string, x int, y int) int {
	var found = 0

	// UP
	if tryXMAS(lines, x, y, 1, 0) {
		found++
	}
	// Down
	if tryXMAS(lines, x, y, -1, 0) {
		found++
	}
	// Forward
	if tryXMAS(lines, x, y, 0, 1) {
		found++
	}
	// Backward
	if tryXMAS(lines, x, y, 0, -1) {
		found++
	}

	// Up Left
	if tryXMAS(lines, x, y, -1, -1) {
		found++
	}
	// Up Right
	if tryXMAS(lines, x, y, -1, 1) {
		found++
	}
	// Down Left
	if tryXMAS(lines, x, y, 1, -1) {
		found++
	}
	// Down Right
	if tryXMAS(lines, x, y, 1, 1) {
		found++
	}

	return found
}

func tryXMAS(lines []string, startX int, startY int, deltaX int, deltaY int) bool {
	// 'X' is at [x,y]
	// Search for 'M' at [x+deltaX, y+deltaY]
	var nextX, nextY = startX + deltaX, startY + deltaY
	if checkChar(lines, nextX, nextY, 'M') {
		nextX, nextY = nextX+deltaX, nextY+deltaY
		if checkChar(lines, nextX, nextY, 'A') {
			nextX, nextY = nextX+deltaX, nextY+deltaY
			if checkChar(lines, nextX, nextY, 'S') {
				return true
			}
		}
	}
	return false
}

func checkChar(lines []string, x int, y int, char uint8) bool {
	if x >= 0 && x < maxX && y >= 0 && y < maxY {
		nextChar := lines[x][y]
		if nextChar == char {
			return true
		} else {
			return false
		}
	}
	return false
}

func day4part2(lines []string) {
	defer timeTrack(time.Now(), "Part2")

	var totalFound = 0

	maxX = len(lines[0])
	maxY = len(lines)

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			char := lines[x][y]
			if char == 'A' {
				totalFound += checkMASSpot(lines, x, y)
			}
		}
	}

	println(totalFound)
}

func checkMASSpot(lines []string, x int, y int) int {
	if x >= 1 && x < maxX-1 && y >= 1 && y < maxY-1 {
		topLeft, topRight := lines[x-1][y-1], lines[x-1][y+1]
		botLeft, botRight := lines[x+1][y-1], lines[x+1][y+1]

		if (topLeft == 'M' && botRight == 'S') || (topLeft == 'S' && botRight == 'M') {
			if (botLeft == 'M' && topRight == 'S') || (botLeft == 'S' && topRight == 'M') {
				return 1
			}
		}
	}
	return 0
}
