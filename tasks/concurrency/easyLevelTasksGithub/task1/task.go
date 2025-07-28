package main

import (
	"fmt"
	"sync"
)

// time.Sleep использовать нельзя. это будет не валидным ответом на собеседовании
// 1
// Что выведет код и почему?
// Как исправить?

//func main() {
//	var counter int
//	for i := 0; i < 1000; i++ {
//		go func() {
//			counter++
//		}()
//	}
//	fmt.Println(counter)
//}

//	1 Вывод будет не ожидаемым/корректным,
//	2 главная горутина не ждет всех горутин (Добавляем горутины в группу и ждем их)
//	3 проблемы в data race, решаем с использование mutex

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	var counter int
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()

			counter++
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
