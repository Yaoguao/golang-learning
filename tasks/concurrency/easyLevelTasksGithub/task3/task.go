package main

import (
	"fmt"
	"sync"
	"time"
)

// 3
// Расскажи подробно что происходит?(спойлер: почему будет panic?)
// Как сделать так, чтобы работало?

//func main() {
//	x := make(map[int]int, 1)
//	go func() { x[1] = 2 }()
//	go func() { x[3] = 7 }()
//	go func() { x[123] = 10 }()
//	go func() { x[1] = 2 }()
//	go func() { x[34] = 7 }()
//	go func() { x[1432] = 10 }()
//	go func() { x[1] = 2 }()
//	go func() { x[100] = 7 }()
//	go func() { x[34] = 10 }()
//	go func() { x[1] = 2 }()
//	time.Sleep(100 * time.Millisecond) //блокируемся на 100 миллисекунд
//	fmt.Println("x[1] =", x[1])
//}

//	Запись в мапу, не потоко-безопастна
//	Два пути
//	-	использовать мапу из пакета sync
//	-	использовать обычную map, но так же использовать и mutex

func main() {
	//x := make(map[int]int, 1)
	m := sync.Map{}
	go func() { m.Store(1, 2) }()
	go func() { m.Store(3, 7) }()
	go func() { m.Store(123, 10) }()
	go func() { m.Store(1, 34) }()
	go func() { m.Store(34, 7) }()
	go func() { m.Store(2323, 23) }()
	go func() { m.Store(2, 7) }()
	go func() { m.Store(3, 745) }()
	go func() { m.Store(34, 7324) }()
	go func() { m.Store(1, 74) }()
	time.Sleep(100 * time.Millisecond) //блокируемся на 100 миллисекунд
	r, _ := m.Load(1)
	fmt.Println("m[1] =", r)
}

//func main() {
//	var mu sync.Mutex
//	x := make(map[int]int, 1)
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[1] = 2
//	}()
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[3] = 7
//	}()
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[123] = 10
//	}()
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[1] = 2
//	}()
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[34] = 7
//	}()
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[1432] = 10
//	}()
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[1] = 2
//	}()
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[100] = 7
//	}()
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[34] = 10
//	}()
//	go func() {
//		mu.Lock()
//		defer mu.Unlock()
//		x[1] = 2
//	}()
//	time.Sleep(100 * time.Millisecond) //блокируемся на 100 миллисекунд
//	fmt.Println("x[1] =", x[1])
//}
