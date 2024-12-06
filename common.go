package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	day6()
}
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func dayX() {
	println("Solving AoC 2024 Day X:")
	input := loadInputDayX()
	dayXpart1(input)
	dayXpart2(input)
}

func loadInputDayX() []string {
	defer timeTrack(time.Now(), "Loading/Parsing Input")
	file, err := os.Open("inputs/dayX.txt")
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

func dayXpart1(lines []string) {
	defer timeTrack(time.Now(), "Part1")

}

func dayXpart2(lines []string) {
	defer timeTrack(time.Now(), "Part2")

}
