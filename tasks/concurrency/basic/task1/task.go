package main

import (
	"fmt"
	"sync"
	"time"
)

//	Напиши программу, которая запускает 5 горутин.
//	Каждая выводит сообщение и "спит" 1 секунду.
//	Главная функция должна дождаться всех

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println("Message GO")
			time.Sleep(1 * time.Second)
		}()
	}

	wg.Wait()

	fmt.Println("Program completion")
}

//func greet(c chan string) {
//	fmt.Println("Hello " + <-c + "!")
//}
//
//func main() {
//	fmt.Println("main() started")
//	c := make(chan string)
//
//	go greet(c)
//
//	c <- "John"
//	fmt.Println("main() stopped")
//}
