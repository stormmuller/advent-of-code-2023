package utils

func MakeFilledMap[T any](size int, defaultValue T) map[int]T {

	var newMap = make(map[int]T)

	for i := 1; i <= size; i++ {
		newMap[i] = defaultValue
	}

	return newMap
}
