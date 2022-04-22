package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создание группы контроля исполниеия горутин 
	var wg sync.WaitGroup
	// Массив
	arr := []int{2, 4, 6, 8, 10}

	// увеличиваем на длинну переменной arr
	wg.Add(len(arr))
	// Запускаем 5 горутин
	for _, e := range arr {
		//способ запуска горутины в анонимной функции
		go func(wg *sync.WaitGroup, e int) {
			fmt.Println(e * e)
			//после завершения исполниия горутины мы уменьшим значение группы на 1
			defer wg.Done()
		}(&wg, e)
	}
	//будем ждать пока в группе не будет нуля
	wg.Wait()
	fmt.Println("Ops... End")
}
