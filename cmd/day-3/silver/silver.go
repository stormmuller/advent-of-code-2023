package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/stormmuller/advent-of-code-2023/utils"
)

func main() {
	args := os.Args[1:] // os.Args[0] is the program path, so skip it

	results := utils.ReadFile(args[0], func(line string) string { return line })

	numbers := getNumbers(results)

	fmt.Println(numbers)

	total := utils.SumArray(numbers)

	fmt.Println(total)
}

func getNumbers(lines []string) []int {
	var numbers []int
	lastLineIndex := len(lines) - 1

	for lineIndex, line := range lines {

		var numbersInLine []int
		var numberBuilder strings.Builder
		isNumberAdjecentToSymbol := false

		for charIndex, char := range line {
			isNumber := unicode.IsDigit(char)

			if isNumber {
				numberBuilder.WriteRune(char)

				if !isNumberAdjecentToSymbol {
					nextLineIndex := lineIndex + 1
					previousLineIndex := lineIndex - 1
					var nextLine string
					var previousLine string

					if nextLineIndex <= lastLineIndex {
						nextLine = lines[nextLineIndex]
					}

					if previousLineIndex >= 0 {
						previousLine = lines[previousLineIndex]
					}

					isNumberAdjecentToSymbol = hasAdjecentSymbols(line, nextLine, previousLine, charIndex)

				}
			} else {
				str := numberBuilder.String()

				if str == "" {
					continue
				}

				numberBuilder.Reset()

				if !isNumberAdjecentToSymbol {
					fmt.Printf("skipping...[%v, %v] %v \n", lineIndex, charIndex, str)
					continue
				}

				isNumberAdjecentToSymbol = false
				number, err := strconv.Atoi(str)

				if err != nil {
					log.Fatal(err)
				}

				// fmt.Printf("The adjecent value for number %v \n", str)

				numbersInLine = append(numbersInLine, number)
			}
		}

		str := numberBuilder.String()

		if str == "" || !isNumberAdjecentToSymbol {
			fmt.Printf("skipping...[%v, ~] %v \n", lineIndex, str)

			numbers = append(numbers, numbersInLine...)
			continue
		}

		number, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}

		numbersInLine = append(numbersInLine, number)

		numbers = append(numbers, numbersInLine...)
	}

	return numbers
}

func hasAdjecentSymbols(currentLine string, nextLine string, previousLine string, charIndex int) bool {

	canCheckLeft := charIndex >= 1
	canCheckRight := charIndex <= len(currentLine)-2 // minus 2 because we will be adding 1
	canCheckDown := nextLine != ""
	canCheckUp := previousLine != ""

	if canCheckLeft {
		previousChar := []rune(currentLine)[charIndex-1]

		if isSpecialCharacter(previousChar) {
			return true
		}
	}

	if canCheckRight {
		nextChar := []rune(currentLine)[charIndex+1]

		if isSpecialCharacter(nextChar) {
			return true
		}
	}

	if canCheckDown {
		charBelow := []rune(nextLine)[charIndex]

		if isSpecialCharacter(charBelow) {
			return true
		}

		if canCheckLeft {
			previousBelowChar := []rune(nextLine)[charIndex-1]

			if isSpecialCharacter(previousBelowChar) {
				return true
			}
		}

		if canCheckRight {
			nextBelowChar := []rune(nextLine)[charIndex+1]

			if isSpecialCharacter(nextBelowChar) {
				return true
			}
		}
	}

	if canCheckUp {
		charAbove := []rune(previousLine)[charIndex]

		if isSpecialCharacter(charAbove) {
			return true
		}

		if canCheckLeft {
			previousAboveChar := []rune(previousLine)[charIndex-1]

			if isSpecialCharacter(previousAboveChar) {
				return true
			}
		}

		if canCheckRight {
			nextAboveChar := []rune(previousLine)[charIndex+1]

			if isSpecialCharacter(nextAboveChar) {
				return true
			}
		}
	}

	return false
}

func isSpecialCharacter(r rune) bool {
	switch r {
	case '-', '*', '@', '=', '%', '/', '+', '&', '^', '#', '$':
		return true
	default:
		return false
	}
}
