package main

import (
	"fmt"
	"sync"
)

//	Напиши пример с гонкой данных (data race), затем исправь его с помощью Mutex.

type MockDB struct {
	//	Добавим mutex в поле
	mu       sync.Mutex
	accounts map[string]int
}

func (db *MockDB) Transfer(from, to string, amount int) {
	//	Используем mutex
	db.mu.Lock()
	defer db.mu.Unlock()

	db.accounts[from] -= amount
	db.accounts[to] += amount
}

func main() {
	db := &MockDB{
		accounts: map[string]int{
			"Alice": 1000,
			"Bob":   1000,
		},
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			db.Transfer("Alice", "Bob", 10)
		}()
	}

	wg.Wait()
	fmt.Println("Final balances:")
	fmt.Println("Alice:", db.accounts["Alice"])
	fmt.Println("Bob:", db.accounts["Bob"])
}
