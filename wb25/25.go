package main

import (
	"fmt"
	"time"
)

// type Duration int64

// func Sleep(d Duration) {}

func Sleep(t time.Duration) {
	time.After(t) // После ожидает истечения длительности, а затем отправляет текущее время
}

func main() {
	fmt.Println("Sleep")

	Sleep(time.Second)

	fmt.Println("two sleep")
}
