package utils

import (
	"bufio"
	"log"
	"os"
)

func Data(filePrefix string, f func(v string) interface{}) chan interface{} {
	c := make(chan interface{})

	go func() {
		defer close(c)

		file, err := os.Open(filePrefix + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- f(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return c
}
