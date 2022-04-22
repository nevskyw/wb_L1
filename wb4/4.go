package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Worker(mainChan chan interface{}, n int, wg *sync.WaitGroup) {
	for data := range mainChan {
		time.Sleep(time.Second)
		fmt.Printf("Worker #%d doing %v\n", n, data)
	}
	wg.Done() //
}

func main() {
	tasks := [...]interface{}{"aaaaa", "bbbbb", 'Q', true} // массив типа interface{}
	mainChan := make(chan interface{})                     // создаем канал типа interface{} куда мы передаем любой список задач

	wg := &sync.WaitGroup{}            // создаем группу ожидания
	sigChan := make(chan os.Signal, 1) // буф
	
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	//quitChan := make(chan struct{})

	rand.Seed(time.Now().UnixNano())

	fmt.Print("Nu of workers: ")
	var n int
	fmt.Scan(&n)

	for n != 0 {
		wg.Add(1)
		go Worker(mainChan, n, wg)
		n--
	}

	for {
		select {
		default:
			mainChan <- tasks[rand.Intn(len(tasks))]
		case <-sigChan:
			close(mainChan)
			wg.Wait()
			return //os.Exit(0)
		}
	}
	// fmt.Println(time.Now().UnixNano())
}
