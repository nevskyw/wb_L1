package main

import "fmt"

func main() {
	var A int
	var B int
	var D float64

	_, aErr := fmt.Scan(&A) // проверяем на ошибку и выводим число, так же вызваем функцию обработчика ошибок
	checkError(aErr)
	_, bErr := fmt.Scan(&B)
	checkError(bErr)

	C := A * B
	fmt.Println("A * B", C) // Выводим значение умножения

	C = A + B
	fmt.Println("A + B", C) // Выводим значение сложения

	C = A - B
	fmt.Println("A - B", C) // Выводим значение разности

	C = A % B
	fmt.Println("A % B", C) // Выводим значение вычисления остатка от деления
	
	C = A / B
	fmt.Println("A / B", C) // Выводим значение целочисленного деления
	
	D = float64(A) / float64(B)
	fmt.Println("A / B", D) // Выводим значение деления чисел с плавающей точкой
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
