package main

import (
	"fmt"
	"os"

	"github.com/stormmuller/advent-of-code-2023/utils"
)

func main() {
	args := os.Args[1:] // os.Args[0] is the program path, so skip it

	results := utils.ReadFile(args[0], processLine)

	total := utils.SumArray(results)

	fmt.Println(total)
}

func processLine(line string) int {

}
