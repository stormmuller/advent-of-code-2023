package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/stormmuller/advent-of-code-2023/utils"
)

func main() {
	args := os.Args[1:] // os.Args[0] is the program path, so skip it

	results := utils.ProcessFile(args[0], processLine)

	total := utils.SumArray(results)

	fmt.Println(total)
}

func processLine(line string) int {
	cleanedLine := cleanLine(line)

	result, err := strconv.Atoi(cleanedLine)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

var numbersAsLetters = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var numbersAsStrings = map[int]string{
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
}

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func cleanLine(line string) string {
	cleanedLine := line

	earliestIndex := 99999
	latestIndex := 0

	var firstNumber, lastNumber int

	for _, number := range numbers {
		numberAsLetters := numbersAsLetters[number]
		numberAsString := numbersAsStrings[number]

		numberAsLettersFirstIndex := strings.Index(cleanedLine, numberAsLetters)
		numberAsStringFirstIndex := strings.Index(cleanedLine, numberAsString)

		numberAsLettersLastIndex := strings.LastIndex(cleanedLine, numberAsLetters)
		numberAsStringLastIndex := strings.LastIndex(cleanedLine, numberAsString)

		if numberAsLettersFirstIndex >= 0 && numberAsLettersFirstIndex <= earliestIndex {
			earliestIndex = numberAsLettersFirstIndex
			firstNumber = number
		}

		if numberAsStringFirstIndex >= 0 && numberAsStringFirstIndex <= earliestIndex {
			earliestIndex = numberAsStringFirstIndex
			firstNumber = number
		}

		if numberAsLettersLastIndex >= 0 && numberAsLettersLastIndex >= latestIndex {
			latestIndex = numberAsLettersLastIndex
			lastNumber = number
		}

		if numberAsStringLastIndex >= 0 && numberAsStringLastIndex >= latestIndex {
			latestIndex = numberAsStringLastIndex
			lastNumber = number
		}
	}

	concatenated := fmt.Sprint(firstNumber) + fmt.Sprint(lastNumber)

	return concatenated
}
