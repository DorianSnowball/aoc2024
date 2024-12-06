package main

import (
	"bufio"
	"os"
	"slices"
	"sort"
	"time"
)

//var field [][]rune

func day6() {
	println("Solving AoC 2024 Day 6:")
	field, guard := loadInputDay6()

	playField := deepCopy(field)

	day6part1(playField, guard)
	day6part2(field, guard, possibleObstacleSpots)
}

func deepCopy(field [][]rune) [][]rune {
	playField := make([][]rune, len(field))
	for i := range field {
		playField[i] = make([]rune, len(field[i]))
		copy(playField[i], field[i])
	}
	return playField
}

func loadInputDay6() ([][]rune, Guard) {
	defer timeTrack(time.Now(), "Loading/Parsing Input")
	file, err := os.Open("inputs/day6.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	lines := make([][]rune, 0)

	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, []rune(line))
	}

	guard := Guard{
		x:         0,
		y:         0,
		direction: []int{0, -1},
	}

	for y, line := range lines {
		var guardPosX = slices.Index(line, '^')
		if guardPosX != -1 {
			guard.x = guardPosX
			guard.y = y
			println("Found Guard Starting Pos at ", guardPosX, y)
		}
	}

	return lines, guard
}

type Guard struct {
	x         int
	y         int
	direction []int
}

var possibleObstacleSpots = make([][]int, 0)

func day6part1(field [][]rune, guard Guard) {
	defer timeTrack(time.Now(), "Part1")

	for step(field, &guard) >= 0 {
		// step until left playing field
	}

	var visitedFields = 0
	for _, bytes := range field {
		for _, b := range bytes {
			if b == '.'+1 || b == '.'+2 || b == '.'+3 {
				visitedFields++
			}
		}
	}

	println(visitedFields)
}

func step(field [][]rune, guard *Guard) rune {
	// Mark currently on field as visited
	if field[guard.y][guard.x] == '^' {
		field[guard.y][guard.x] = '.'
	}

	field[guard.y][guard.x] += 1
	visited := field[guard.y][guard.x] - '.'

	frontPos := []int{guard.x + guard.direction[0], guard.y + guard.direction[1]}
	if frontPos[0] < 0 || frontPos[0] >= len(field) || frontPos[1] < 0 || frontPos[1] >= len(field[0]) {
		// Stepped outside!
		//println("Stepped outside!")
		return -1
	}

	frontChar := field[frontPos[1]][frontPos[0]]
	if frontChar == '#' {
		rotate(guard)
		//println("Rotated at", guard.x, guard.y, "to", guard.direction[0], guard.direction[1])
	} else {
		possibleObstacleSpots = append(possibleObstacleSpots, frontPos)
	}

	guard.x += guard.direction[0]
	guard.y += guard.direction[1]

	return visited
}

func rotate(guard *Guard) {
	guard.direction = []int{-guard.direction[1], guard.direction[0]}
	/*
		0, -1 Up
		1, 0 Right
		0, 1 Down
		-1, 0 Left
	*/
}

func day6part2(field [][]rune, guard Guard, spots [][]int) {
	defer timeTrack(time.Now(), "Part2")

	println("Possible Obstacle Spots:", len(spots))
	// Remove duplicate possible spots
	slices.SortFunc(spots, func(a, b []int) int {
		if a[1] == b[1] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})
	spots = slices.CompactFunc(spots, func(a []int, b []int) bool {
		return a[0] == b[0] && a[1] == b[1]
	})

	println("Dedublicated Obstacle Spots:", len(spots))

	startPos := []int{guard.x, guard.y}
	index := slices.IndexFunc(spots, func(pos []int) bool {
		return pos[0] == startPos[0] && pos[1] == startPos[1]
	})

	if index != -1 {
		spots = append(spots[:index], spots[index+1:]...)
		println("Removed Start Pos from possible array")
	}

	detectedSpots := make([][]int, 0)
	maxVisited := make([]int, 0)

	for _, spot := range spots {
		//println("Checking next Spot:", spot[1], spot[0])
		// clone field to have an independent playfield
		playField := deepCopy(field)
		playField[spot[1]][spot[0]] = '#'

		visited := step(playField, &guard)
		for visited >= 0 {
			// step until left playing field or loop detected
			if visited > 3 {
				// Detected loop!
				detectedSpots = append(detectedSpots, spot)
				//maxVisited = append(maxVisited, getMaxVisited(playField))
				break
			}

			visited = step(playField, &guard)
		}

		if spot[0] == 6 || spot[1] == 7 {
		}
		// reset guard for next possible spot
		guard.x = startPos[0]
		guard.y = startPos[1]
		guard.direction = []int{0, -1}
	}

	sort.Ints(maxVisited)
	//println("Max Visited:", maxVisited[len(maxVisited)-1])

	println(len(detectedSpots))
}

func getMaxVisited(field [][]rune) int {
	maxVisited := 0
	for i, runes := range field {
		println(string(runes))
		for j, _ := range runes {
			char := field[i][j]
			var visited int
			visited = int(char - '.')
			if visited > maxVisited {
				maxVisited = visited
			}
		}
	}
	return maxVisited
}
