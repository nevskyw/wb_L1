package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mux         sync.Mutex
	information map[string]string
}

func main() {
	mux := sync.RWMutex{}

	newMap := make(map[int]int)

	ch := make(chan int)

	for i := 1; i < 6; i++ {
		go func(i int) {
			mux.Lock()
			newMap[i] = i * i // записываем по очереди
			mux.Unlock()
			ch <- i
		}(i)
	}

	for i := 0; i < 5; i++ {
		// получаем из канала ключ по которому лежит значение
		x := <-ch // вытаскиваем по эти ключам
		fmt.Println("Ключ: ", x, "Значение: ", newMap[x])
	}
}
