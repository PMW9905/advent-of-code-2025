package main

import (
	"fmt"
	"parker/aoc2025/util"
	"strconv"
)

const StartingLocation = 50
const LockLength = 100

func getRotation(line string) (byte, int, error) {
	direction := line[0]
	clicksAsString := line[1:]

	clicks, err := strconv.Atoi(clicksAsString)

	return direction, clicks, err
}

func partOne(lines []string) (int, error) {
	password := 0
	location := StartingLocation

	for _, line := range lines {
		direction, clicks, err := getRotation(line)

		if err != nil {
			return -1, err
		}

		adjustedClicks := clicks % LockLength

		if direction == 'L' {
			adjustedClicks = LockLength - adjustedClicks
		}

		location = (location + adjustedClicks) % LockLength


		if location == 0 {
			password += 1
		}

		fmt.Println(direction, clicks, location, password)
	}

	return password, nil
}

func main() {
	lines, err := utils.ReadInput("day-01/input.txt")

	if err != nil {
		panic(err)
	}

	partOneAnswer, err := partOne(lines)
	fmt.Printf("Part 1 answer: %d\n", partOneAnswer)
}
