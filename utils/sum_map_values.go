package utils

func SumMapValues[TKey comparable, TValue Number](mapToSum map[TKey]TValue) TValue {
	var total TValue = 0

	for _, number := range mapToSum {
		total += number
	}

	return total
}
