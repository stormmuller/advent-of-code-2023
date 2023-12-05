package utils

import (
	"bufio"
	"log"
	"os"
)

type textHandler[T any] func(string) T

func ProcessFile[T any](path string, textHandler textHandler[T]) []T {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var results []T

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		result := textHandler(line)
		results = append(results, result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return results
}

func ReadFile(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var results []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		results = append(results, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return results
}
