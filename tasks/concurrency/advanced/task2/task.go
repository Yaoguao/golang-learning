package main

import (
	"fmt"
	"sync"
	"time"
)

//	Сделай воркер-пул: N воркеров читают из канала задачи и обрабатывают их.
//	Главная горутина пишет в канал 100 заданий.
//	Используй WaitGroup, чтобы дождаться выполнения.

func worker(tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range tasks {
		time.Sleep(time.Millisecond)

		fmt.Printf("LOG: %d\n", t)
	}
}

func main() {
	var wg sync.WaitGroup
	tasks := make(chan int)
	var n int
	fmt.Scanln(&n)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(tasks, &wg)
	}

	for i := 1; i <= 100; i++ {
		tasks <- i
	}
	close(tasks)
	fmt.Println("CLOSE")
	fmt.Println("wrote 100 tasks")

	//3
	wg.Wait()

}
