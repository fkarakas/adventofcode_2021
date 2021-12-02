package main

import (
	"github.com/fkarakas/adventofcode_2021/utils"
)

func Part1() int {
	previousDepth := 0
	increaseCount := 0

	for i, depth := range utils.LoadDataAsInt() {
		if i == 0 {
			previousDepth = depth
			continue
		}
		if depth > previousDepth {
			increaseCount++
		}
		previousDepth = depth
	}

	return increaseCount
}

func Part2() int {
	currentWindowDepth := 0
	previousWindowDepth := 0
	increaseCount := 0

	arr := utils.LoadDataAsInt()

	// init previous window
	if len(arr) > 3 {
		previousWindowDepth = arr[0] + arr[1] + arr[2]
	}

	for i := 1; i < len(arr); i++ {
		if i+2 >= len(arr) {
			break
		}

		currentWindowDepth = arr[i] + arr[i+1] + arr[i+2]

		if currentWindowDepth > previousWindowDepth {
			increaseCount++
		}
		previousWindowDepth = currentWindowDepth
	}

	return increaseCount
}

func main() {
	println("part1 increase count=", Part1())
	println("part2 increase count=", Part2())
}
