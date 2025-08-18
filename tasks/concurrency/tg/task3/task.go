package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Как сделать влияние функции getDiscount() на всю программу более контролируемым?

// Эта функция лезет по сети в старый монолит и может тупить.
// TODO - или использовать context.WithTimeout
// TODO - можно вернуть error
func getDiscount(timeout time.Duration) float64 {
	resCh := make(chan float64)

	go func() {
		time.Sleep(time.Duration(rand.Int63n(2)) * time.Second)

		resCh <- 12.0
		//close(resCh)
	}()

	select {
	case <-time.After(timeout):
		return 0
	case v := <-resCh:
		return v
	}
}

func main() {
	fmt.Printf("Ваша скидка: %v", getDiscount(time.Second))
}
