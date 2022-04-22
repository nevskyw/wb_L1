package main

import "fmt"

// бинарный поиск ищет определенное значение
// https://www.golangprograms.com/golang-program-for-implementation-of-binary-search.html

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	fmt.Println(binarySearch(63, items)) // Число и слайс в котором будем искать
}

func binarySearch(num int, newSlice []int) bool {

	low := 0                  // начало среза
	high := len(newSlice) - 1 // конец среза

	for low <= high {
		// 4
		centre := (low + high) / 2 // Высчитываем середину

		if newSlice[centre] < num {
			low = centre + 1
		} else {
			high = centre - 1
		}
	}

	// Если ушли в начало и ничего не нашли || Если ушли в конец ничего не нашли,
	// то равно false
	if low == len(newSlice) || newSlice[low] != num {
		return false
	}

	return true
}
