package main

import (
	"fmt"
	"time"
)

// Что выведет программа и почему?

//	Думал что будет deadlock, но у меня почему-то вечное ожидание, что странно
//	Вечное ожидание, из-за блокировки главной горутины и доп. гор

func main() {
	ch := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Отдельная горутина отвисла")
		ch <- false
	}()

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Произошел тик тикера")
			ch <- true
		case value := <-ch:
			fmt.Printf("Получено значение %t\n", value)
			return
		}
	}
}
