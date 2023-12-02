package utils

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func SumArray[T Number](numbers []T) T {
	var total T = 0

	for _, number := range numbers {
		total += number
	}

	return total
}
