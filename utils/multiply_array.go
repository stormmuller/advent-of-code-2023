package utils

func MultiplyArray[T Number](numbers []T) T {
	var total T = 1

	for _, number := range numbers {
		total *= number
	}

	return total
}
