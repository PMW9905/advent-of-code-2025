package main

import (
	"fmt"
	"parker/aoc2025/util"
)

const MostPaper = 3

func isInBounds(r int, c int, lines [][]byte) bool {
	return r >= 0 && r < len(lines) && c >= 0 && c < len(lines[0])
}

func locationValue(r int, c int, lines [][]byte) int {
	if isInBounds(r, c, lines) && lines[r][c] == '@' {
		return 1
	}
	return 0
}

func surroundingPaper(r int, c int, lines [][]byte) int {
	sum := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			sum += locationValue(r+i, c+j, lines)
		}
	}

	return sum
}

func surroundingPaperDFS(r int, c int, lines [][]byte) int {
	numRemoved := 0
	sum := surroundingPaper(r, c, lines)

	if sum <= MostPaper {
		numRemoved = 1
		lines[r][c] = '.'
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				if isInBounds(r+i, c+j, lines) && lines[r+i][c+j] == '@' {
					numRemoved += surroundingPaperDFS(r+i, c+j, lines)
				}
			}
		}
	}
	return numRemoved
}

func partTwo(lines [][]byte) int {
	reachablePaper := 0
	for r := range lines {
		for c := range lines[0] {
			if lines[r][c] == '.' {
				continue
			}
			reachablePaper += surroundingPaperDFS(r, c, lines)
		}
	}
	return reachablePaper
}

func partOne(lines [][]byte) int {
	reachablePaper := 0
	for r := range lines {
		for c := range lines[0] {
			if lines[r][c] == '.' {
				continue
			}
			surrounding := surroundingPaper(r, c, lines)
			if surrounding <= MostPaper {
				reachablePaper += 1
				fmt.Println(r, c, surrounding)
			}
		}
	}
	return reachablePaper
}

func main() {
	lines, err := utils.ReadInput("day-04/input.txt")

	linesbytes := make([][]byte, len(lines))
	for i := range lines {
		linesbytes[i] = []byte(lines[i])
	}

	if err != nil {
		panic(err)
	}

	partOneAnswer := partOne(linesbytes)

	fmt.Printf("Part 1 answer: %d\n", partOneAnswer)

	partTwoAnswer := partTwo(linesbytes)

	fmt.Printf("Part 2 answer: %d\n", partTwoAnswer)
}
