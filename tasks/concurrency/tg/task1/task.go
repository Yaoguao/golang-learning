package main

import (
	"fmt"
	"sync"
)

// Напищите функцию GetOrCreate, которая создает новый элемент мапы,
// если его еше не было и возвращает его значение, или просто возвращает значение при наличии.
// Важно учесть, что код должен нормально работать в конкурентной среде.

type MapConc struct {
	values map[string]string
	mu     *sync.Mutex
}

func NewConcurrentMap() MapConc {
	return MapConc{
		values: make(map[string]string),
		mu:     &sync.Mutex{},
	}
}

func (m *MapConc) GetOrCreate(key, val string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	if v, ok := m.values[val]; ok {
		return v
	}

	m.values[key] = val
	return val
}

func main() {
	cm := NewConcurrentMap()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		val := cm.GetOrCreate("key1", "value1")
		fmt.Println("Goroutine 1 got:", val)
	}()

	go func() {
		defer wg.Done()
		val := cm.GetOrCreate("key1", "value2")
		fmt.Println("Goroutine 2 got:", val)
	}()

	wg.Wait()
}
