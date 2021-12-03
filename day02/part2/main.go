package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/fkarakas/adventofcode_2021/utils"
)

func commands() chan interface{} {
	return utils.Data("../input", func(v string) interface{} {
		arr := strings.Split(v, " ")

		value, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatal(err)
		}

		return command{
			direction: arr[0],
			value:     value,
		}

	})
}

type command struct {
	direction string
	value     int
}

type submarine struct {
	aim  int
	x, y int
}

func (s *submarine) Move(c command) {
	switch c.direction {
	case "forward":
		s.Forward(c.value)
	case "down":
		s.Down(c.value)
	case "up":
		s.Up(c.value)
	}
}

func (s *submarine) Forward(v int) {
	s.x += v
	s.y += s.aim * v
}

func (s *submarine) Down(v int) {
	s.aim += v
}

func (s *submarine) Up(v int) {
	s.aim -= v
}

func response2() int {
	sub := submarine{}
	for c := range commands() {
		cmd := c.(command)
		sub.Move(cmd)
	}
	return sub.x * sub.y
}

func main() {
	fmt.Printf("response2=%v\n", response2())
}
