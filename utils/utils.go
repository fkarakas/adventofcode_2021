package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func LoadDataAsInt() []int {
	result := []int{}
	for _, s := range LoadData() {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, i)
	}
	return result
}

func LoadData() []string {
	result := []string{}

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
