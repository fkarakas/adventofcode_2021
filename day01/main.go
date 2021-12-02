package main

import (
	"github.com/fkarakas/adventofcode_2021/utils"
)

func Part1() int {
	previousDepth := -1
	increaseCount := 0

	for depth := range utils.Depths() {
		if previousDepth == -1 {
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
	buffer := []int{}
	increaseCount := 0

	for depth := range utils.Depths() {
		if len(buffer) < 4 {
			buffer = append(buffer, depth)
			continue
		}

		currentWindow := buffer[0] + buffer[1] + buffer[2]
		nextWindow := buffer[1] + buffer[2] + buffer[3]

		if nextWindow > currentWindow {
			increaseCount++
		}

		buffer = append(buffer[1:], depth)
	}

	currentWindow := buffer[0] + buffer[1] + buffer[2]
	nextWindow := buffer[1] + buffer[2] + buffer[3]

	if nextWindow > currentWindow {
		increaseCount++
	}

	return increaseCount
}

func main() {
	println("part1 increase count=", Part1())
	println("part2 increase count=", Part2())
}
