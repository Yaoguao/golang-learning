package main

import (
	"fmt"
	"sync"
)

// Задание 18
// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговоезначение счетчика.
// устраняем гонку данных
type Counter struct {
	mx    sync.RWMutex
	Count int
}

func (c *Counter) Increment() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.Count++
}
func (c *Counter) Decrement() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.Count--
}

func main() {
	wg := sync.WaitGroup{}
	c := Counter{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}
	wg.Wait()

	fmt.Println(c.Count)
}
