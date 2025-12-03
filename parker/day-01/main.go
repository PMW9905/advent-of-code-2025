package main

import (
	"fmt"
	"parker/aoc2025/util"
	"strconv"
)

const TotalDialNumbers = 100
const StartingDialPosition = 50

<<<<<<< HEAD
func parseRotationAsRight(rotation string) (int, error) {
	direction := rotation[0]
	numStr := rotation[1:]
=======
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
>>>>>>> 9a743dcd77c8196b1a17c6a1a4311cc7fd8deab1

	num, err := strconv.Atoi(numStr)

	if direction == 'L' {
		num = 100 - num
	}
	return num, err
}

func countStopsOnZero(turns []string) int {
	password := 0
	position := StartingDialPosition

<<<<<<< HEAD
	for _, rotation := range turns {
		num, err := parseRotationAsRight(rotation)
=======
	for _, line := range lines {
		clicks, err := getSignedRotation(line)
>>>>>>> 9a743dcd77c8196b1a17c6a1a4311cc7fd8deab1

		if err != nil {
			panic(err)
		}

<<<<<<< HEAD
		position += num
		position = (position) % TotalDialNumbers

		if position == 0 {
			password += 1
		}
	}
	return password
=======
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
>>>>>>> 9a743dcd77c8196b1a17c6a1a4311cc7fd8deab1
}

func main() {
	lines, err := utils.ReadInput("day-01/input.txt")

	if err != nil {
		panic(err)
	}

	password := countStopsOnZero(lines)
	fmt.Println("Part 1 Answer:", password)
}
