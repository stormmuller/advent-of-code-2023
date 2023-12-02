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

	results := utils.ReadFile(args[0], processLine)

	total := utils.SumArray(results)

	fmt.Println(total)
}

const MAX_RED_CUBES int = 12
const MAX_GREEN_CUBES int = 13
const MAX_BLUE_CUBES int = 14

func processLine(line string) int {
	splitString := strings.Split(line, ":")
	gameNumber := strings.Split(splitString[0], " ")[1]
	rounds := strings.Split(splitString[1], ";")

	maxColors := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, round := range rounds {
		turns := strings.Split(round, ",")
		for _, turnString := range turns {
			turn := strings.Split(strings.TrimSpace(turnString), " ")
			colorAmount, err := strconv.Atoi(turn[0])

			if err != nil {
				log.Fatal(err)
			}

			color := turn[1]

			currentColorAmount := maxColors[color]

			if colorAmount > currentColorAmount {
				maxColors[color] = colorAmount
			}
		}
	}

	isPossibleGame := maxColors["red"] <= MAX_RED_CUBES &&
		maxColors["green"] <= MAX_GREEN_CUBES &&
		maxColors["blue"] <= MAX_BLUE_CUBES

	if !isPossibleGame {
		return 0
	}

	result, err := strconv.Atoi(gameNumber)

	if err != nil {
		log.Fatal(err)
	}

	return result
}
