package main

import (
	"bufio"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func day1() {
	list1, list2 := loadInputDay1()
	day1part1(list1, list2)
	day1part2(list1, list2)
}

func loadInputDay1() ([]int, []int) {
	defer timeTrack(time.Now(), "Loading/Parsing Input")
	file, err := os.Open("inputs/day1.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for scanner.Scan() {
		var line = scanner.Text()
		split := strings.Split(line, "   ")
		if len(split) == 2 {
			var val1, _ = strconv.Atoi(split[0])
			var val2, _ = strconv.Atoi(split[1])

			list1 = append(list1, val1)
			list2 = append(list2, val2)
		}
	}
	return list1, list2
}

func day1part1(list1, list2 []int) {
	defer timeTrack(time.Now(), "Part1")
	sort.Ints(list1)
	sort.Ints(list2)

	var totalDifference = 0
	for i := 0; i < len(list1); i++ {
		lineDifference := list1[i] - list2[i]
		totalDifference += max(lineDifference, -lineDifference)
	}

	println(totalDifference)
}

func day1part2(list1, list2 []int) {
	defer timeTrack(time.Now(), "Part2")
	var similarityScore = 0
	for i := 0; i < len(list1); i++ {
		var n = list1[i]
		position, found := slices.BinarySearch(list2, n)
		if !found {
			continue
		}

		var count = 0
		for list2[position] == n {
			count++
			position++
		}
		similarityScore += n * count
	}

	println(similarityScore)
}
