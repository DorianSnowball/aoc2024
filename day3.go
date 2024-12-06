package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"time"
)

func day3() {
	println("Solving AoC 2024 Day 3:")
	lines := loadInputDay3()
	day3part1(lines)
	day3part2(lines)
}

func loadInputDay3() []string {
	defer timeTrack(time.Now(), "Loading/Parsing Input")
	file, err := os.Open("inputs/day3.txt")
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

func day3part1(lines []string) {
	defer timeTrack(time.Now(), "Part1")

	var mulPattern = regexp.MustCompile("(mul\\(\\d{1,3},\\d{1,3}\\))")

	var sum = 0
	for _, line := range lines {
		mulExpressions := mulPattern.FindAllString(line, -1)
		for _, expression := range mulExpressions {
			var numberPattern = regexp.MustCompile("(\\d{1,3})")
			matchedNumbers := numberPattern.FindAllString(expression, -1)
			if len(matchedNumbers) != 2 {
				panic("Matched Mul did not include 2 numbers")
			}
			i, _ := strconv.Atoi(matchedNumbers[0])
			j, _ := strconv.Atoi(matchedNumbers[1])

			sum += i * j
		}
	}

	println(sum)
}

func day3part2(lines []string) {
	defer timeTrack(time.Now(), "Part2")

	//var line = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	var matchPattern = regexp.MustCompile("(?:mul\\(\\d{1,3},\\d{1,3}\\))|(?:don't\\(\\))|(?:do\\(\\))")

	var sum = 0
	var enabled = true
	for _, line := range lines {

		matchExpressions := matchPattern.FindAllString(line, -1)
		for _, expression := range matchExpressions {

			if expression == "do()" {
				enabled = true
				continue
			} else if expression == "don't()" {
				enabled = false
				continue
			}
			if !enabled {
				continue
			}
			var numberPattern = regexp.MustCompile("(\\d{1,3})")
			matchedNumbers := numberPattern.FindAllString(expression, -1)
			if len(matchedNumbers) != 2 {
				panic("Matched Mul did not include 2 numbers")
			}
			i, _ := strconv.Atoi(matchedNumbers[0])
			j, _ := strconv.Atoi(matchedNumbers[1])

			sum += i * j
		}
	}

	println(sum)
}
