package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func foo(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	d := rand.Intn(401) + 100
	time.Sleep(time.Duration(d) * time.Millisecond)
	fmt.Printf("successfully: %d\n", n)
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go foo(i, &wg)
	}

	wg.Wait()
	fmt.Println("CLOSE")
}
