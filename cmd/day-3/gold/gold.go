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

	results := utils.ProcessFile(args[0], func(line string) string { return line })

	numbers := getNumbers(results)

	total := utils.SumArray(numbers)

	fmt.Println(total)
}

func getNumbers(lines []string) []int {
	var numbers []int

	for lineIndex, line := range lines {
		for charIndex, char := range line {
			if isGear(char) {
				adjacentNumbers := getAdjecentNumbers(lines, lineIndex, charIndex)

				if len(adjacentNumbers) == 2 {
					total := 1

					for _, num := range adjacentNumbers {
						total *= num
					}

					numbers = append(numbers, total)
				}
			}
		}
	}

	return numbers
}

func getAdjecentNumbers(lines []string, lineIndex int, charIndex int) map[coordinate]int {
	canCheckUp := lineIndex > 0
	canCheckRight := charIndex <= len(lines[lineIndex])-2 // minus 2 because we will be adding 1
	canCheckDown := lineIndex < len(lines)-1
	canCheckLeft := charIndex > 0

	var nChar, neChar, eChar, seChar, sChar, swChar, wChar, nwChar rune

	if canCheckUp {
		nChar = []rune(lines[lineIndex-1])[charIndex]
	}

	if canCheckUp && canCheckRight {
		neChar = []rune(lines[lineIndex-1])[charIndex+1]
	}

	if canCheckRight {
		eChar = []rune(lines[lineIndex])[charIndex+1]
	}

	if canCheckDown && canCheckRight {
		seChar = []rune(lines[lineIndex+1])[charIndex+1]
	}

	if canCheckDown {
		sChar = []rune(lines[lineIndex+1])[charIndex]
	}

	if canCheckDown && canCheckLeft {
		swChar = []rune(lines[lineIndex+1])[charIndex-1]
	}

	if canCheckLeft {
		wChar = []rune(lines[lineIndex])[charIndex-1]
	}

	if canCheckUp && canCheckLeft {
		nwChar = []rune(lines[lineIndex-1])[charIndex-1]
	}

	// var numbers []int
	numbers := map[coordinate]int{}

	// nChar, neChar, eChar, seChar, sChar, swChar, wChar, nwChar

	if unicode.IsDigit(nChar) {
		coord, number := battleShipNumber(lines, lineIndex-1, charIndex)
		numbers[coord] = number
	}

	if unicode.IsDigit(neChar) {
		coord, number := battleShipNumber(lines, lineIndex-1, charIndex+1)
		numbers[coord] = number
	}

	if unicode.IsDigit(eChar) {
		coord, number := battleShipNumber(lines, lineIndex, charIndex+1)
		numbers[coord] = number
	}

	if unicode.IsDigit(seChar) {
		coord, number := battleShipNumber(lines, lineIndex+1, charIndex+1)
		numbers[coord] = number
	}

	if unicode.IsDigit(sChar) {
		coord, number := battleShipNumber(lines, lineIndex+1, charIndex)
		numbers[coord] = number
	}

	if unicode.IsDigit(swChar) {
		coord, number := battleShipNumber(lines, lineIndex+1, charIndex-1)
		numbers[coord] = number
	}

	if unicode.IsDigit(wChar) {
		coord, number := battleShipNumber(lines, lineIndex, charIndex-1)
		numbers[coord] = number
	}

	if unicode.IsDigit(nwChar) {
		coord, number := battleShipNumber(lines, lineIndex-1, charIndex-1)
		numbers[coord] = number
	}

	return numbers
}

func battleShipNumber(lines []string, lineIndex int, startIndex int) (coordinate, int) {
	line := lines[lineIndex]

	start := calculateNumberStart(line, startIndex)

	var sb strings.Builder

	var currentIndex int = start

	for currentIndex < len(line) {
		char := rune(line[currentIndex])

		if !unicode.IsDigit(char) {
			break
		}

		sb.WriteRune(char)
		currentIndex++
	}

	result, err := strconv.Atoi(sb.String())

	if err != nil {
		log.Fatal(err)
	}

	return coordinate{lineIndex, start}, result
}

func calculateNumberStart(line string, index int) int {
	if index == 0 {
		return index
	}

	if unicode.IsDigit(rune(line[index-1])) {
		return calculateNumberStart(line, index-1)
	}

	return index
}

func isGear(r rune) bool {
	if r == '*' {
		return true
	}

	return false
}

type coordinate struct {
	line  int
	index int
}
