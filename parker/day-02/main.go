package main

import (
	"fmt"
	"parker/aoc2025/util"
	"strconv"
	"strings"
)

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

func calculateLeftFloor(left int, leftStr string) int {

	if len(leftStr) <= 1 {
		return 1
	}

	halfSize := len(leftStr) / 2
	// check if odd.
	// if odd, return pow(10, len(leftStr) / 2) 12,345 -> 100, 123 -> 10, etc.
	if len(leftStr)%2 == 1 {
		return utils.Pow(10, halfSize)
	}
	// if even, split in half.
	splitter := utils.Pow(10, halfSize)
	firstHalf := left / splitter  // 1234 -> 12
	secondHalf := left % splitter // 1234 -> 34
	// if first half >= second half, return first half.
	if firstHalf >= secondHalf {
		return firstHalf
	} else {
		// otherwise, return first half + 1, since
		// value firstHalf can't be in the range.
		return firstHalf + 1
	}
}

func calculateRightCeil(right int, rightStr string) int {

	halfSize := len(rightStr) / 2
	// check if odd.
	// if odd, return pow(10, len(rightStr) / 2) - 1:  12,345 -> 99, 123 -> 9, etc.
	if len(rightStr)%2 == 1 {
		return utils.Pow(10, halfSize) - 1
	}
	// if even, split in half.
	splitter := utils.Pow(10, halfSize)
	firstHalf := right / splitter  // 1234 -> 12
	secondHalf := right % splitter // 1234 -> 34
	// if first half >= second half, return first half - 1, since
	// value firsthalffirsthalf can't be in the range
	if firstHalf > secondHalf {
		return firstHalf - 1
	} else {
		// otherwise, return first half
		return firstHalf
	}
}

func sumFakeIDs(left int, leftStr string, right int, rightStr string) int {
	leftFloor := calculateLeftFloor(left, leftStr)
	rightCeil := calculateRightCeil(right, rightStr)
	fmt.Println(leftFloor)
	fmt.Println(rightCeil)

	fakeIDSum := 0
	multiplier := utils.Pow(10, len(leftStr)/2)
	for i := leftFloor; i <= rightCeil; i++ {
		if multiplier <= i {
			multiplier *= 10
		}
		fakeIDSum += i*multiplier + i
	}
	return fakeIDSum
}

func partOne(idRangesString string) int {
	idRangesSlice := strings.Split(idRangesString, ",")
	fakeIDs := 0

	for _, idRange := range idRangesSlice {
		fakeIDs += sumFakeIDs(getIDRange(idRange))
		fmt.Println("RUNNING TALLY: ", fakeIDs)
	}
	return fakeIDs
}

func fuknBruteForceBaby(left int, leftStr string, right int, rightStr string) int {
}

func partTwo(idRangesString string) int {
	idRangesSlice := strings.Split(idRangesString, ",")
	fakeIDs := 0

	for _, idRange := range idRangesSlice {
		fakeIDs += fuknBruteForceBaby(getIDRange(idRange))
		fmt.Println("RUNNING TALLY: ", fakeIDs)
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

	partTwoAnswer := partTwo(idRangesString)

	fmt.Printf("Part 2 answer: %d\n", partTwoAnswer)
}
