package main

import (
	"bufio"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func day5() {
	println("Solving AoC 2024 Day 5:")
	rules, updates := loadInputDay5()
	invalidUpdates := day5part1(rules, updates)
	day5part2(rules, invalidUpdates)
}

func loadInputDay5() ([][]int, [][]int) {
	defer timeTrack(time.Now(), "Loading/Parsing Input")
	file, err := os.Open("inputs/day5.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	rules := make([][]int, 0)
	updates := make([][]int, 0)

	for scanner.Scan() {
		var line = scanner.Text()
		if strings.Contains(line, "|") {
			// Parsing Rule
			split := strings.Split(line, `|`)
			constrained, _ := strconv.Atoi(split[0])
			before, _ := strconv.Atoi(split[1])
			rules = append(rules, []int{constrained, before})
		} else if strings.Contains(line, ",") {
			// Parsing update
			split := strings.Split(line, `,`)
			pages := make([]int, 0)
			for _, s := range split {
				page, _ := strconv.Atoi(s)
				pages = append(pages, page)
			}

			updates = append(updates, pages)
		}
	}
	return rules, updates
}

func day5part1(rules [][]int, updates [][]int) [][]int {
	defer timeTrack(time.Now(), "Part1")

	var middlePageSum = 0

	invalidUpdates := make([][]int, 0)

	// build PageRuleMap for ease of use
	rulesMap := createRulesMap(rules)

	for _, update := range updates {
		// Build ignored rules for update
		ignoredPages := getIgnoredPages(update, rulesMap)
		// Check order of update
		// For each page, check if conform to rules
		valid := checkValidity(update, rulesMap, ignoredPages)
		if valid {
			middlePageSum += update[len(update)/2]
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	println(middlePageSum)
	return invalidUpdates
}

func createRulesMap(rules [][]int) map[int][]int {
	rulesMap := make(map[int][]int)
	for _, rule := range rules {
		constrained := rule[0]
		before := rule[1]
		rulesMap[constrained] = append(rulesMap[constrained], before)
	}
	return rulesMap
}

func getIgnoredPages(update []int, rulesMap map[int][]int) []int {
	ignoredPages := make([]int, 0)
	for key := range maps.Keys(rulesMap) {
		if !slices.Contains(update, key) {
			ignoredPages = append(ignoredPages, key)
		}
	}

	for valueSlice := range maps.Values(rulesMap) {
		for _, value := range valueSlice {
			if !slices.Contains(update, value) {
				if !slices.Contains(ignoredPages, value) {
					ignoredPages = append(ignoredPages, value)
				}
			}
		}
	}
	return ignoredPages
}

func checkValidity(update []int, rulesMap map[int][]int, ignoredPages []int) bool {
	valid := true
	for i, page := range update {
		constraints := rulesMap[page]

		after := update[i:]
		for _, constraint := range constraints {
			if !slices.Contains(ignoredPages, constraint) && !slices.Contains(after, constraint) {
				valid = false
			}
		}

	}
	return valid
}

func day5part2(rules [][]int, invalidUpdates [][]int) {
	defer timeTrack(time.Now(), "Part2")

	var middlePageSum = 0

	rulesMap := createRulesMap(rules)

	stillInvalid, middlePageSum := fixUpdates(invalidUpdates, rulesMap, middlePageSum)

	for len(stillInvalid) > 0 {
		stillInvalid, middlePageSum = fixUpdates(stillInvalid, rulesMap, middlePageSum)
	}
	println(middlePageSum)
}

func fixUpdates(invalidUpdates [][]int, rulesMap map[int][]int, middlePageSum int) ([][]int, int) {
	stillInvalid := make([][]int, 0)
	for _, update := range invalidUpdates {
		ignoredPages := getIgnoredPages(update, rulesMap)

		for i, page := range update {
			constraints := rulesMap[page]

			after := update[i+1:]
			for _, constraint := range constraints {
				if !slices.Contains(ignoredPages, constraint) && !slices.Contains(after, constraint) {
					// swap places with colliding page
					collidingPageIdx := slices.Index(update, constraint)
					update[i], update[collidingPageIdx] = update[collidingPageIdx], update[i]

				}
			}
		}

		// Check order of update
		// For each page, check if conform to rules
		valid := checkValidity(update, rulesMap, ignoredPages)
		if valid {
			middlePageSum += update[len(update)/2]
		} else {
			stillInvalid = append(stillInvalid, update)
		}
	}
	return stillInvalid, middlePageSum
}
