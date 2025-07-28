package main

import (
	"fmt"
	"sync"
)

// Напиши counter с использованием sync.Mutex, увеличивающий значение из 1000 горутин.

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
