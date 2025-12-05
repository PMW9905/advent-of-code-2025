package main

import (
	"fmt"
	"parker/aoc2025/util"
	"strconv"
	"strings"
)

func pow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

func getIDRange(idRangeString string) (int, string, int, string) {
	split := strings.Split(idRangeString, "-")
	left, err := strconv.Atoi(split[0])

	if err != nil {
		panic(err)
	}

	right, err := strconv.Atoi(split[1])

	if err != nil {
		panic(err)
	}

	fmt.Println(left, split[0], right, split[1])

	return left, split[0], right, split[1]
}

func calculateRangeEdge(isFloor bool, val int, valStr string) int {
	halfLength := len(valStr) / 2

	floorOffset := 0
	ceilOffset := 0

	if isFloor {
		floorOffset = 1
	}
	ceilOffset = 1 - floorOffset

	if len(valStr)%2 == 1 {
		return pow(10, halfLength+1) - ceilOffset
	}

	firstHalf := val / pow(10, halfLength-1)
	lastHalf := val % pow(10, halfLength-1)

	if firstHalf >= lastHalf {
		return firstHalf - ceilOffset
	} else {
		return lastHalf + ceilOffset
	}
}

func countFakeIDs(left int, leftStr string, right int, rightStr string) int {
	leftmostID := calculateRangeEdge(true, left, leftStr)
	rightmostID := calculateRangeEdge(false, right, rightStr) 
	fmt.Println(leftmostID)
	fmt.Println(rightmostID)
	return rightmostID - leftmostID
}

func partOne(idRangesString string) int {
	idRangesSlice := strings.Split(idRangesString, ",")
	fakeIDs := 0

	for _, idRange := range idRangesSlice {
		fakeIDs += countFakeIDs(getIDRange(idRange))
	}
	return fakeIDs
}

func main() {
	lines, err := utils.ReadInput("day-02/input_small.txt")

	if err != nil {
		panic(err)
	}

	idRangesString := lines[0]

	partOneAnswer := partOne(idRangesString)

	fmt.Printf("Part 1 answer: %d\n", partOneAnswer)
}
