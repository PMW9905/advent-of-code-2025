package main

import (
	"fmt"
	"parker/aoc2025/util"
	"strconv"
)

const StartingLocation = 50
const LockLength = 100

func getSignedRotation(line string) (int, error) {
	direction := line[0]
	clicksAsString := line[1:]

	clicks, err := strconv.Atoi(clicksAsString)

	if direction == 'L' {
		clicks *= -1
	}

	return clicks, err
}

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
		clicks, err := getSignedRotation(line)

		if err != nil {
			return -1, err
		}

		location = ((clicks % LockLength) + location) % LockLength

		if location == 0 {
			password += 1
		}
	}

	return password, nil
}

func partTwo(lines []string) (int, error) {
	password := 0
	location := StartingLocation

	for _, line := range lines {
		direction, clicks, err := getRotation(line)

		if err != nil {
			return -1, err
		}

		password += clicks / LockLength

		adjustedClicks := clicks % LockLength

		if direction == 'L' {
			if location-adjustedClicks <= 0 && location != 0 {
				password += 1
			}
			adjustedClicks = LockLength - adjustedClicks
		} else {
			if location+adjustedClicks >= LockLength {
				password += 1
			}
		}

		location = (location + adjustedClicks) % LockLength
	}

	return password, nil
}

func main() {
	lines, err := utils.ReadInput("day-01/input.txt")

	if err != nil {
		panic(err)
	}

	partOneAnswer, err := partOne(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1 answer: %d\n", partOneAnswer)

	partTwoAnswer, err := partTwo(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 2 answer: %d\n", partTwoAnswer)

}
