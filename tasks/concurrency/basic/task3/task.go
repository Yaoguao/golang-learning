package main

import (
	"fmt"
	"sync"
)

//	Сделай простую "фабрику": одна горутина кладёт задачи (числа) в канал, другая читает и удваивает число.

func main() {
	num := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(num)

		for i := 1; i <= 10; i++ {
			num <- i
		}
	}()

	go func() {
		defer wg.Done()

		for val := range num {
			//	Можно записывать в другой канал результат
			fmt.Println(val * val)
		}
	}()

	wg.Wait()
}
