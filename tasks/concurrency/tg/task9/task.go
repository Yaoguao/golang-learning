package main

import "sync"

// Напишите реализацию InMemory кэша

type ICache interface {
	Set(k string, v string)
	Get(k string) (string, bool)
}

type InMemory struct {
	cache map[string]string
	mu    *sync.RWMutex
}

func NewCache(cap int) InMemory {
	return InMemory{
		cache: make(map[string]string, cap),
		mu:    &sync.RWMutex{},
	}

}

func (c *InMemory) Set(k string, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[k] = v
}

func (c *InMemory) Get(k string) (string, bool) {
	c.mu.RLock()
	val, ok := c.cache[k]
	c.mu.RUnlock()

	return val, ok
}

//	Попытка создать LRU кеш

//type LRUCache struct {
//	cache map[string]*list.Element
//	mu    *sync.Mutex
//	list  *list.List
//}
//
//func NewLRUCache(cap int) LRUCache {
//	return LRUCache{
//		cache: make(map[string]*list.Element, cap),
//		mu:    &sync.Mutex{},
//		list:  list.New(),
//	}
//}
//
//func (c *LRUCache) Get(k string) (string, bool) {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//
//	if elem, ok := c.cache[k]; ok {
//		c.list.MoveToFront(elem)
//		return elem.Value.(string), true
//	}
//	return "", false
//}
//
//func (c *LRUCache) Set(k string, v string) {
//	c.mu.Lock()
//	defer c.mu.Unlock()
//
//	if elem, ok := c.cache[k]; ok {
//		c.list.MoveToFront(elem)
//		elem.Value = v
//		return
//	}
//
//	if
//}

func main() {

}
