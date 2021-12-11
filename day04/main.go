package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fkarakas/adventofcode_2021/utils"
)

const (
	file = "data.txt"
)

func getNumbers() chan int {
	c := make(chan int)

	go func() {
		defer close(c)

		for line := range utils.Data(file, nil) {
			for _, s := range strings.Split(line.(string), ",") {
				n, err := strconv.Atoi(s)
				if err == nil {
					c <- n
				}
			}
			break
		}
	}()

	return c
}

type board struct {
	data [][]int
}

func (b *board) set(line, column, value int) {
	if line >= len(b.data) {
		nbLines := line - len(b.data)
		for i := 0; i <= nbLines; i++ {
			b.data = append(b.data, []int{})
		}
	}
	if column >= len(b.data[line]) {
		nbColumns := column - len(b.data[line])
		for i := 0; i <= nbColumns; i++ {
			b.data[line] = append(b.data[line], 0)
		}
	}
	b.data[line][column] = value
}

func (b board) get(line, col int) int {
	return b.data[line][col]
}

func (b board) lasIndex(line int) int {
	return len(b.data[line]) - 1
}

func (b board) numberOfLines() int {
	return len(b.data)
}

func (b board) numberOfColumns(line int) int {
	return len(b.data[line])
}

func (b board) String() string {
	var sb strings.Builder
	for l := 0; l < b.numberOfLines(); l++ {
		for c := 0; c < b.numberOfColumns(l); c++ {
			sb.WriteString(fmt.Sprintf("%v", b.get(l, c)))
			if c != b.lasIndex(l) {
				sb.WriteString("\t")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (b board) mark(n int) {
	for l := 0; l < b.numberOfLines(); l++ {
		for c := 0; c < b.numberOfColumns(l); c++ {
			if b.get(l, c) == n {
				b.set(l, c, -1)
			}
		}
	}
}

func (b board) hasWon() bool {
	// check lines
	countMark := 0
	for l := 0; l < b.numberOfLines(); l++ {
		if b.get(l, 0) != -1 {
			continue
		}
		countMark = 0
		for c := 0; c < b.numberOfColumns(l); c++ {
			if b.get(l, c) == -1 {
				countMark++
			}
		}
		if countMark == b.numberOfColumns(l) {
			return true
		}
	}
	// check columns
	countMark = 0
	for c := 0; c < b.numberOfColumns(0); c++ {
		if b.get(0, c) != -1 {
			continue
		}
		countMark = 0
		for l := 0; l < b.numberOfLines(); l++ {
			if b.get(l, c) == -1 {
				countMark++
			}
		}
		if countMark == b.numberOfLines() {
			return true
		}
	}
	return false
}

func (b board) sumAllUnmarked() int {
	sum := 0
	for l := 0; l < b.numberOfLines(); l++ {
		for c := 0; c < b.numberOfColumns(l); c++ {
			if b.get(l, c) != -1 {
				sum += b.get(l, c)
			}
		}
	}
	return sum
}

func getBoards() chan board {
	c := make(chan board)

	go func() {
		defer close(c)

		skipFirst := true

		var b board
		line := -1

		for lineData := range utils.Data(file, nil) {
			if skipFirst {
				skipFirst = false
				continue
			}
			if lineData == "" {
				if line != -1 {
					c <- b
				}
				b = board{}
				line = 0
				continue
			}
			for i, s := range strings.Fields(lineData.(string)) {
				n, err := strconv.Atoi(s)
				if err == nil {
					b.set(line, i, n)
				}
			}
			line++
		}
		c <- b
	}()

	return c
}

func isin(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// part1 response=63552 (sample 4512)
// part2 response=9020 (sample 1924)
func main() {
	boards := []board{}
	for b := range getBoards() {
		boards = append(boards, b)
		fmt.Printf("%v\n\n", b)
	}
	winners := []int{}
loop:
	for nb := range getNumbers() {
		fmt.Printf("nb=%v ", nb)
		for i, b := range boards {
			if isin(winners, i) {
				continue
			}
			b.mark(nb)
			if b.hasWon() {
				winners = append(winners, i)
				fmt.Printf("\nwinning board: %v", i)
				fmt.Printf(" sum=%v score=%v\n", b.sumAllUnmarked(), b.sumAllUnmarked()*nb)
			}
			if len(winners) == len(boards) {
				fmt.Printf("Done...")
				break loop
			}
		}
	}
	fmt.Printf("\n\n=============\n")
	for _, b := range boards {
		fmt.Printf("%v\n", b)
	}
	/*
		fmt.Printf("\nfirst winning board=%v\n", winningBoard)
		for _, b := range boards {
			fmt.Printf("%v\n\n", b)
		}
		fmt.Printf("sum=%v\n", boards[winningBoard].sumAllUnmarked())
		fmt.Printf("final score=%v\n", boards[winningBoard].sumAllUnmarked()*lastNumber)
	*/
}
