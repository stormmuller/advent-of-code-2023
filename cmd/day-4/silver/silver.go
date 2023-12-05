package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/stormmuller/advent-of-code-2023/utils"
)

func main() {
	args := os.Args[1:] // os.Args[0] is the program path, so skip it

	results := utils.ReadFile(args[0], processCard)

	fmt.Println(results)

	total := utils.SumArray(results)

	fmt.Println(total)
}

func processCard(line string) int {
	results := strings.Split(line, "|")

	played := results[0]
	winning := results[1]

	playedNumbers := strings.Split(played, ":")[1]
	playedNumbersArray := utils.RemoveEmptyStrings(strings.Split(playedNumbers, " "))
	winningNumbersArray := utils.RemoveEmptyStrings(strings.Split(winning, " "))

	numberOfMatches := 0

	for _, winningNumber := range winningNumbersArray {
		for _, playedNumber := range playedNumbersArray {
			if winningNumber == playedNumber {
				numberOfMatches++
			}
		}
	}

	return doubleXTimes(numberOfMatches)
}

func doubleXTimes(x int) int {
	if x == 0 {
		return 0
	}

	result := 1
	for i := 0; i < x-1; i++ {
		result *= 2
	}
	return result
}
