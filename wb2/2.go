package main

import "fmt"

var array = []int{2, 4, 6, 8, 10}

// функция squares
func Squares(c chan int) {
	// перебираем - перемножаем и отправляем в канал значения
	for _, value := range array {
		c <- value * value
	}

	// закрываем канал
	close(c)
}

func main() {

	fmt.Printf("Запустить массив: %v\n", array)

	// создаем канал, который принимает только целочисленные значения
	// запускаем горутину
	c := make(chan int)
	go Squares(c)

	for v := range c {
		fmt.Println("Значение в квадрате: ", v)
	}
}
