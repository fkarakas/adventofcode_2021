package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fkarakas/adventofcode_2021/utils"
)

func depths() chan interface{} {
	return utils.Data("input", func(v string) interface{} {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		return i
	})
}

func response1() int {
	previousDepth := -1
	increaseCount := 0

	for d := range depths() {
		depth := d.(int)
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

func calculateIncrease(buffer []int, increaseCount int) int {
	currentWindow := buffer[0] + buffer[1] + buffer[2]
	nextWindow := buffer[1] + buffer[2] + buffer[3]

	if nextWindow > currentWindow {
		increaseCount++
	}
	return increaseCount
}

func response2() int {
	// a buffer of 4 to compare current and next window
	buffer := []int{}
	increaseCount := 0

	for d := range depths() {
		depth := d.(int)

		if len(buffer) < 4 {
			buffer = append(buffer, depth)
			continue
		}

		increaseCount = calculateIncrease(buffer, increaseCount)

		buffer = append(buffer[1:], depth)
	}

	increaseCount = calculateIncrease(buffer, increaseCount)

	return increaseCount
}

func main() {
	fmt.Printf("response1=%v\n", response1())
	fmt.Printf("response2=%v\n", response2())
}
