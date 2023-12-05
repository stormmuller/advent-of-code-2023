package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/stormmuller/advent-of-code-2023/utils"
)

var cardCount = map[int]int{}

func main() {
	args := os.Args[1:] // os.Args[0] is the program path, so skip it

	results := utils.ReadFile(args[0])

	cardCount = utils.MakeFilledMap(len(results), 0)

	for cardNumber := range cardCount {
		processCards(results, cardNumber)

		fmt.Printf("processed %v\n", cardNumber)
	}

	total := utils.SumMapValues(cardCount)

	fmt.Println(total)
}

func processCards(cards []string, cardNumber int) {
	index := cardNumber - 1

	cardCount[cardNumber]++

	numberOfMatches := getNumberOfMatches(cards[index])

	for i := cardNumber + 1; i < cardNumber+numberOfMatches+1; i++ {
		processCards(cards, i)
	}
}

func getNumberOfMatches(card string) int {
	results := strings.Split(card, "|")

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

	return numberOfMatches
}
