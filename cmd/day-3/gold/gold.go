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

	total := utils.SumArray(numbers)

	fmt.Println(total)
}

func getNumbers(lines []string) []int {
	var numbers []int
	lastLineIndex := len(lines) - 1
	gearsMap := map[string][]int{}

	for lineIndex, line := range lines {

		var numbersInLine []int
		var numberBuilder strings.Builder
		isNumberAdjecentToGear := false
		var gearsAdjacentToNumber []string

		for charIndex, char := range line {
			isNumber := unicode.IsDigit(char)

			if isNumber {
				numberBuilder.WriteRune(char)

				if !isNumberAdjecentToGear {
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

					gearsAdjacentToNumber = append(gearsAdjacentToNumber, getAdjecentGears(line, nextLine, previousLine, charIndex, lineIndex)...)

					isNumberAdjecentToGear = len(gearsAdjacentToNumber) > 0
				}
			} else {
				str := numberBuilder.String()

				if str == "" {
					continue
				}

				numberBuilder.Reset()

				if !isNumberAdjecentToGear {
					continue
				}

				isNumberAdjecentToGear = false
				number, err := strconv.Atoi(str)

				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(gearsAdjacentToNumber)

				for _, gear := range gearsAdjacentToNumber {
					gearsMap[gear] = append(gearsMap[gear], number)
				}

				numbersInLine = append(numbersInLine, number)
			}
		}

		str := numberBuilder.String()

		if str == "" || !isNumberAdjecentToGear {
			numbers = append(numbers, numbersInLine...)
			continue
		}

		number, err := strconv.Atoi(str)

		if err != nil {
			log.Fatal(err)
		}

		for _, gear := range gearsAdjacentToNumber {
			gearsMap[gear] = append(gearsMap[gear], number)
		}

		numbersInLine = append(numbersInLine, number)

		numbers = append(numbers, numbersInLine...)
	}

	fmt.Println(gearsMap)

	return numbers
}

func getAdjecentGears(currentLine string, nextLine string, previousLine string, charIndex int, lineIndex int) []string {

	canCheckLeft := charIndex >= 1
	canCheckRight := charIndex <= len(currentLine)-2 // minus 2 because we will be adding 1
	canCheckDown := nextLine != ""
	canCheckUp := previousLine != ""

	var gears []string

	if canCheckLeft {
		previousChar := []rune(currentLine)[charIndex-1]

		if isGear(previousChar) {
			gear := strings.Join([]string{strconv.Itoa(lineIndex), strconv.Itoa(charIndex - 1)}, "")
			gears = append(gears, gear)
		}
	}

	if canCheckRight {
		nextChar := []rune(currentLine)[charIndex+1]

		if isGear(nextChar) {
			gear := strings.Join([]string{strconv.Itoa(lineIndex), strconv.Itoa(charIndex + 1)}, "")
			gears = append(gears, gear)
		}
	}

	if canCheckDown {
		charBelow := []rune(nextLine)[charIndex]

		if isGear(charBelow) {
			gear := strings.Join([]string{strconv.Itoa(lineIndex + 1), strconv.Itoa(charIndex)}, "")
			gears = append(gears, gear)
		}

		if canCheckLeft {
			previousBelowChar := []rune(nextLine)[charIndex-1]

			if isGear(previousBelowChar) {
				gear := strings.Join([]string{strconv.Itoa(lineIndex + 1), strconv.Itoa(charIndex - 1)}, "")
				gears = append(gears, gear)
			}
		}

		if canCheckRight {
			nextBelowChar := []rune(nextLine)[charIndex+1]

			if isGear(nextBelowChar) {
				gear := strings.Join([]string{strconv.Itoa(lineIndex + 1), strconv.Itoa(charIndex + 1)}, "")
				gears = append(gears, gear)
			}
		}
	}

	if canCheckUp {
		charAbove := []rune(previousLine)[charIndex]

		if isGear(charAbove) {
			gear := strings.Join([]string{strconv.Itoa(lineIndex - 1), strconv.Itoa(charIndex)}, "")
			gears = append(gears, gear)
		}

		if canCheckLeft {
			previousAboveChar := []rune(previousLine)[charIndex-1]

			if isGear(previousAboveChar) {
				gear := strings.Join([]string{strconv.Itoa(lineIndex - 1), strconv.Itoa(charIndex - 1)}, "")
				gears = append(gears, gear)
			}
		}

		if canCheckRight {
			nextAboveChar := []rune(previousLine)[charIndex+1]

			if isGear(nextAboveChar) {
				gear := strings.Join([]string{strconv.Itoa(lineIndex - 1), strconv.Itoa(charIndex + 1)}, "")
				gears = append(gears, gear)
			}
		}
	}

	return gears
}

func isGear(r rune) bool {
	if r == '*' {
		return true
	}

	return false
}
