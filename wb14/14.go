package main

import (
	"fmt"
	"reflect"
)

func main() {
	var valueInterface interface{} // Объявляем переменную типа интерфейс

	valueInterface = 1
	fmt.Println(reflect.TypeOf(valueInterface)) // Выводим тип переменной интерфейса - int
	valueInterface = "Hi"
	fmt.Println(reflect.TypeOf(valueInterface)) // Выводим тип переменной интерфейса - string
	valueInterface = true
	fmt.Println(reflect.TypeOf(valueInterface)) // Выводим тип переменной интерфейса - bool
	valueInterface = make(chan int)
	fmt.Println(reflect.TypeOf(valueInterface)) // Выводим тип переменной интерфейса - channel
}

/*
// способ 2 через ТИП fmt

type MyType struct{}

func main() {
	var i int
	var bo bool
	var st string

	arr := []interface{}{}
	arr = append(arr, 1, true, "Привет\n", i, bo, st)

	// через ТИП под копотом fmt
	for _, t := range arr {
		fmt.Print(t)
		fmt.Printf(" это тип: %T\n", t)
	}
}

*/
