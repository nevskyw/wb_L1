package main

import (
	"fmt"
	"time"
)

//Timer срабатывает один раз;
//Ticker будет срабатывать снова и снова до вызова метода Stop().
//
// Функция записи в основной канал
func wrt(ch chan int, val int) {
	ch <- val
}

// функция чтения из канал, где инициализировали v и записали данные основного канала
func rd(ch chan int) {
	v := <-ch
	fmt.Println("Received data with value: ", v) // показываем визуальное действие
}

func main() {
	var t time.Duration // переменная типа time.Duration
	fmt.Print("Set time in seconds: ")
	fmt.Scan(&t)                            // сканируем эту переменную
	ch := make(chan int)                    // создаем канал, в который мы будем писать и читать
	timer := time.NewTimer(t * time.Second) //
	i := 1                                  // итерация числа по каналу???

	for {
		select {
		case <-timer.C: // получаем сигнал по истечению времени
			fmt.Println("time is up")
			close(ch) // закрываем основной канал по истечению времени
			return
		default:
			go wrt(ch, i) // запуск горутины записи, куда передаю ch и val
			go rd(ch)     // запуск горутины чтения основного канала
			time.Sleep(100 * time.Millisecond) // Чтобы мы не бежали по числам, как угорелые
			i++ // итерация по времени
		}
	}
}
