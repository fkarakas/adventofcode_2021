package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fkarakas/adventofcode_2021/utils"
)

const (
	file   = "data.txt"
	width  = 1000
	height = 1000
)

func toPoint(s string) point {
	arr := strings.Split(s, ",")
	x, _ := strconv.Atoi(arr[0])
	y, _ := strconv.Atoi(arr[1])

	return point{x: x, y: y}
}

func segments() chan segment {
	c := make(chan segment)

	go func() {
		defer close(c)

		for line := range utils.Data(file, nil) {
			pointsArr := strings.Split(line.(string), " -> ")

			seg := segment{start: toPoint(pointsArr[0]), end: toPoint(pointsArr[1])}

			c <- seg
		}
	}()

	return c
}

type segment struct {
	start, end point
}

type point struct {
	x, y int
}

type matrix struct {
	data [height][width]int
}

func (m matrix) get(x, y int) int {
	return m.data[y][x]
}

func (m matrix) width() int {
	return len(m.data[0])
}

func (m matrix) height() int {
	return len(m.data)
}

func (m *matrix) add(x, y, value int) {
	m.data[y][x] += value
}

func lower(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func higher(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func (m *matrix) drawLine(x1, y1, x2, y2 int, newValue func(value int)) {
	if x1 == x2 {
		//fmt.Printf("vertical line: %v %v %v %v\n", x1, y1, x2, y2)
		for y := lower(y1, y2); y <= higher(y1, y2); y++ {
			m.add(x1, y, 1)
			if newValue != nil {
				newValue(m.get(x1, y))
			}
		}
	} else if y1 == y2 {
		//fmt.Printf("horizontal line: %v %v %v %v\n", x1, y1, x2, y2)
		for x := lower(x1, x2); x <= higher(x1, x2); x++ {
			m.add(x, y1, 1)
			if newValue != nil {
				newValue(m.get(x, y1))
			}
		}
	} else {
		if x2 < x1 {
			x, y := x1, y1
			x1, y1 = x2, y2
			x2, y2 = x, y
		}
		coefficientDirecteur := (y2 - y1) / (x2 - x1)
		//fmt.Printf("cd=%v\n", coefficientDirecteur)
		for x1 != x2 && y1 != y2 {
			//fmt.Printf("x=%v y=%v", x1, y1)
			m.add(x1, y1, 1)
			if newValue != nil {
				newValue(m.get(x1, y1))
			}
			x1++
			y1 += coefficientDirecteur
		}
		m.add(x1, y1, 1)
		if newValue != nil {
			newValue(m.get(x1, y1))
		}
	}
}

func (m matrix) Print() {
	for y := 0; y < m.height(); y++ {
		for x := 0; x < m.width(); x++ {
			fmt.Printf("%v", m.get(x, y))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// part1 : reponse sample is 5
// part1 : response data is 7438
func main() {
	m := matrix{}

	count := 0
	countDangerousZone := func(value int) {
		if value == 2 {
			count++
		}
	}

	for seg := range segments() {
		//fmt.Printf("seg=%v\n", seg)
		m.drawLine(seg.start.x, seg.start.y, seg.end.x, seg.end.y, countDangerousZone)
	}

	//m.drawLine(0, 9, 5, 9)

	//m.Print()
	fmt.Printf("count dangerous zone=%v", count)
}
