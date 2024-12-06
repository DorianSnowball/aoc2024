package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func day2() {
	println("Solving AoC 2024 Day 2:")
	reports := loadInputDay2()
	day2part1(reports)
	day2part2(reports)
}

func loadInputDay2() [][]int {
	defer timeTrack(time.Now(), "Loading/Parsing Input")
	file, err := os.Open("inputs/day2.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	reports := make([][]int, 0)

	for scanner.Scan() {
		var line = scanner.Text()
		split := strings.Split(line, " ")
		var report = make([]int, 0)
		if len(split) > 1 {
			for _, s := range split {
				var i, _ = strconv.Atoi(s)
				report = append(report, i)
			}
		} else {
			panic("Split is len > 1")
		}
		reports = append(reports, report)
	}
	return reports
}

func day2part1(reports [][]int) {
	defer timeTrack(time.Now(), "Part1")

	var safeReports = 0
	for _, report := range reports {
		safe := checkReport(report)
		if safe {
			safeReports++
		}
	}

	println(safeReports)
}

func checkReport(report []int) bool {
	safe := true
	ascending := checkAsc(report)
	for i := 0; i < len(report)-1; i++ {
		current := report[i]
		next := report[i+1]

		var difference = 0
		if ascending {
			difference = next - current
		} else {
			difference = current - next
		}

		if difference < 1 || difference > 3 {
			// This report is NOT safe
			safe = false
		}
	}
	return safe
}

func checkAsc(report []int) bool {
	if report[0] > report[1] {
		return false
	} else {
		return true
	}
}

func day2part2(reports [][]int) {
	defer timeTrack(time.Now(), "Part2")

	var safeReports = 0
	for _, report := range reports {
		safe := checkReport(report)
		if safe {
			safeReports++
		} else {
			for i := 0; i < len(report); i++ {

				var attempt = make([]int, len(report))
				copy(attempt, report)
				attempt = append(attempt[:i], attempt[i+1:]...)
				safe = checkReport(attempt)
				if safe {
					safeReports++
					break
				}
			}
		}
	}

	println(safeReports)
}

//if i+2 < len(report) {
//					report = append(report[:i+1], report[i+2:]...)
//				} else {
//					errors++
//					break
//				}
