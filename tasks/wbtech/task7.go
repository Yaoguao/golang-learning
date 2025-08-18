package main

import (
	"fmt"
	"sync"
)

// Задание 7
// Реализовать конкурентную запись данных в map.
// ** Вариант с sync.Map
type MapSec struct {
	mx sync.RWMutex
	m  map[string]int
}

func (ms *MapSec) Set(key string, value int) {
	ms.mx.Lock()
	defer ms.mx.Unlock()
	ms.m[key] = value
}

func main() {
	m := MapSec{
		mx: sync.RWMutex{},
		m:  make(map[string]int),
	}
	m.Set("Test", 1)
	fmt.Println(m.m)
}
