package main

import "fmt"

/*
Основная идея - берем опорную точку (centre), проходим массив, чтобы элементы,
которые меньше опорной точки оказались слева от нее, а которые больше - справа.
Дальше берем часть массива до опорной точки и вторую часть после опортной точки, повторяем на них сортировку.
Продолжаем до того момента, как сортируемая часть массива будет пустой или состоять из одного элемента.

*/

func main() {
	arr := []int{3, 4, 2, 1, 5, -7, -1, 0}
	Quicksort(arr) // вызываем функцию сортировки 
	fmt.Println(arr)

}

// создаем  функцию сортировки - Quicksort
func Quicksort(arr []int) {
	// если длинна массива меньше либо равна 1 то возвращаем
	if len(arr) <= 1 {
		return
	}

	split := Partition(arr) // Запустили функцию Partition для сортировки массива
	Quicksort(arr[:split])  // СОРТИРОВКА с лево на право
	Quicksort(arr[split:])  // СОРТИРОВКА с право на лево

}

// создаем функцию разделения - Partition
func Partition(arr []int) int {
	centre := arr[len(arr)/2] // определяем середину массива

	left := 0             // начало массива
	right := len(arr) - 1 // конец массива

	// Берем часть массива до опорной точки и вторую часть после опортной точки, повторяем на них сортировку
	// Продолжаем до того момента, как сортируемая часть массива будет пустой или состоять из одного элемента.
	for {
		for ; arr[left] < centre; left++ {
		}
		for ; arr[right] < centre; right++ {
		}

		// свап
		arr[left], arr[right] = arr[right], arr[left]
	}

}

/*
https://xn----7sbaabdq1a3ayardngchc4a.xn--p1ai/%D0%B7%D0%B0%D0%BC%D0%B5%D1%82%D0%BA%D0%B8/%D0%B0%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC%D1%8B-%D1%81%D0%BE%D1%80%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%BA%D0%B8-%D0%BD%D0%B0-go/
http://algolist.manual.ru/sort/quick_sort.php
*/
