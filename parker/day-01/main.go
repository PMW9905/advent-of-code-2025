package main

import (
	"fmt"
	"parker/aoc2025/util"
)

func main() {
	lines, err := utils.ReadInput("day-01/input.txt")

	if err != nil {
		panic(err)
	}
	fmt.Println(lines)
}
