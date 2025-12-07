package main

import (
	"container/list"
	"fmt"
	"parker/aoc2025/util"
)

const MaxNumBatteries = 12

func byteToInt(b byte) int {
	return int(b) - int('0')
}

func partTwo(lines []string) int {
	sum := 0
	head := list.New()

	for _, line := range lines {
		head.Init()

		lastIndex := len(line) - 1
		for i := lastIndex; i >= 0; i-- {
			curDigit := byteToInt(line[i])
			head.PushFront(curDigit)

			if head.Len() <= MaxNumBatteries {
				continue
			}

			for e := head.Front(); e != nil; e = e.Next() {
				if e.Next() == nil || e.Value.(int) < e.Next().Value.(int) {
					head.Remove(e)
					break
				}
			}
		}

		thisSum := 0
		for e := head.Front(); e != nil; e = e.Next() {
			thisSum *= 10
			thisSum += e.Value.(int)
		}
		sum += thisSum
	}
	return sum
}

func partOne(lines []string) int {
	sum := 0
	for _, line := range lines {
		lastIndex := len(line) - 1
		largestDigitSeen := byteToInt(line[lastIndex])
		largestPairSeen := -1

		for i := lastIndex - 1; i >= 0; i-- {
			curDigit := byteToInt(line[i])

			largestPairSeen = max(largestPairSeen, 10*curDigit+largestDigitSeen)
			largestDigitSeen = max(largestDigitSeen, curDigit)
		}
		sum += largestPairSeen
	}
	return sum
}

func main() {
	lines, err := utils.ReadInput("day-03/input.txt")

	if err != nil {
		panic(err)
	}

	partOneAnswer := partOne(lines)

	fmt.Printf("Part 1 answer: %d\n", partOneAnswer)

	partTwoAnswer := partTwo(lines)

	fmt.Printf("Part 2 answer: %d\n", partTwoAnswer)

}
