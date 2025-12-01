package main

import (
	"fmt"
	"parker/aoc2025/util"
	"strconv"
)

const TotalDialNumbers = 100
const StartingDialPosition = 50

func parseRotationAsRight(rotation string) (int, error) {
	direction := rotation[0]
	numStr := rotation[1:]

	num, err := strconv.Atoi(numStr)

	if direction == 'L' {
		num = 100 - num
	}
	return num, err
}

func countStopsOnZero(turns []string) int {
	password := 0
	position := StartingDialPosition

	for _, rotation := range turns {
		num, err := parseRotationAsRight(rotation)

		if err != nil {
			panic(err)
		}

		position += num
		position = (position) % TotalDialNumbers

		if position == 0 {
			password += 1
		}
	}
	return password
}

func main() {
	lines, err := utils.ReadInput("day-01/input.txt")

	if err != nil {
		panic(err)
	}

	password := countStopsOnZero(lines)
	fmt.Println("Part 1 Answer:", password)
}
