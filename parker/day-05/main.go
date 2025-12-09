package main

import (
	"fmt"
	"parker/aoc2025/util"
	"sort"
	"strconv"
	"strings"
)

type NumRange struct {
	Left  int
	Right int
}

func getRangesAndIDs(lines []string) ([]NumRange, []int) {
	numRanges := []NumRange{}
	ids := []int{}
	index := 0
	for lines[index] != "" {
		splitRange := strings.Split(lines[index], "-")
		left, err := strconv.Atoi(splitRange[0])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(splitRange[1])
		if err != nil {
			panic(err)
		}
		numRanges = append(numRanges, NumRange{left, right})
		index += 1
	}
	index += 1
	for index < len(lines) {
		num, err := strconv.Atoi(lines[index])
		if err != nil {
			panic(err)
		}
		ids = append(ids, num)
		index += 1
	}
	return numRanges, ids
}

func partOne(numRanges []NumRange, ids []int) int {
	numFresh := 0

	for _, id := range ids {
		for _, numRange := range numRanges {
			if numRange.Left <= id && id <= numRange.Right {
				numFresh += 1
				break
			}
		}
	}

	return numFresh
}

func partTwo(numRanges []NumRange) int {

	compressedRanges := []NumRange{}

	sort.Slice(numRanges, func(i, j int) bool {
		if numRanges[i].Left != numRanges[j].Left {
			return numRanges[i].Left < numRanges[j].Left
		}
		return numRanges[i].Right < numRanges[j].Right
	})

	compressedRanges = append(compressedRanges, numRanges[0])

	for i := 1; i < len(numRanges); i++ {
		numRange := numRanges[i]

		compressedRight := compressedRanges[len(compressedRanges)-1].Right
		if numRange.Left <= compressedRight {
			compressedRanges[len(compressedRanges)-1].Right = max(compressedRight, numRange.Right)
		} else {
			compressedRanges = append(compressedRanges, numRange)
		}
	}

	numIds := 0
	for _, compressedRange := range compressedRanges {
		numIds += compressedRange.Right - compressedRange.Left + 1
	}

	return numIds
}

func main() {
	lines, err := utils.ReadInput("day-05/input.txt")

	if err != nil {
		panic(err)
	}

	numRanges, ids := getRangesAndIDs(lines)

	partOneAnswer := partOne(numRanges, ids)

	fmt.Printf("Part 1 answer: %d\n", partOneAnswer)

	partTwoAnswer := partTwo(numRanges)

	fmt.Printf("Part 2 answer: %d\n", partTwoAnswer)
}
