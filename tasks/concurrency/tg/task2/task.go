package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Требуется доработать метод SimulateRequest так, чтобы:
// - Код работал в конкурентной среде
// - При долгом ожидании запрос отваливался по таймауту
// - На консоль печаталось время выполнения запроса

var counter int64
var mu sync.Mutex // TODO - Это плохое решение, но я еще не уверен, что можно передавать его как аргумент функции

func SimulateRequest(timeout time.Duration) (int64, error) { // TODO - или использовать timer
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	signalRes := make(chan int64)
	t := time.Now()
	defer func() {
		log.Println("Time request:", time.Since(t))
	}()

	go func() {
		time.Sleep(time.Duration(rand.Int63n(5)) * time.Second)

		mu.Lock() // TODO - Если сильно изменять аргументы функции нельзя, то есть вариант использовать атомики
		counter++
		mu.Unlock()
		signalRes <- counter
	}()

	select {
	case <-ctx.Done():
		return 0, errors.New("request stopped of long wait timeout")
	case v := <-signalRes:
		return v, nil

	}

}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			val, err := SimulateRequest(3 * time.Second)
			if err != nil {
				log.Println("ERROR", err)
				return
			}
			log.Printf("Значение счетчика: %d\n", val)
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}
