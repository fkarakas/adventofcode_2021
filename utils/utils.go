package utils

import (
	"bufio"
	"log"
	"os"
)

func Data(file string, mapFilter func(v string) interface{}) chan interface{} {
	c := make(chan interface{})

	go func() {
		defer close(c)

		file, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			value := mapFilter(scanner.Text())
			if value != nil {
				c <- value
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return c
}
