package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func Depths() chan int {
	c := make(chan int)

	go func() {
		defer close(c)

		file, err := os.Open("data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- toInt(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return c
}
