package main

import (
	"fmt"
	"time"
)

//	Задание 6 *
//	Реализовать все возможные способы остановки выполнения горутины.

// 1 способ использовать канал
func worker(ch chan bool) {
	for {
		select {
		case <-ch:
			fmt.Println("Остановлен")
			return
		default:
			fmt.Println("Работаем")
			time.Sleep(1 * time.Second)
		}

	}
}

// 2 способ контекст?...

func main() {
	ch := make(chan bool)

	go worker(ch)

	time.Sleep(5 * time.Second)
	ch <- true

	time.Sleep(1 * time.Second)
	fmt.Println("Программа завершена")
}
