package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"

	"github.com/stormmuller/advent-of-code-2023/utils"
)

func main() {
	args := os.Args[1:] // os.Args[0] is the program path, so skip it

	results := utils.ProcessFile(args[0], processLine)

	total := utils.SumArray(results)

	fmt.Println(total)
}

func processLine(line string) int {
	var firstNumber, lastNumber rune

	for _, char := range line {
		if unicode.IsDigit(char) {
			if firstNumber == 0 {
				firstNumber = char
			}

			lastNumber = char
		}
	}

	concatenated := string(firstNumber) + string(lastNumber)
	result, err := strconv.Atoi(concatenated)

	if err != nil {
		log.Fatal(err)
	}

	return result
}
