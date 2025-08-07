package main

// time.Sleep использовать нельзя. это будет не валидным ответом на собеседовании
// 1. Иногда приходят нули. В чем проблема? Исправь ее
// 2. Если функция bank_network_call выполняется 5 секунд, то за сколько выполнится balance()?
//	- 25 секунд?

//func balance() int {
//	x := make(map[int]int, 1)
//	var m sync.Mutex
//	wg := sync.WaitGroup{}
//
//	wg.Add(5)
//	// call bank
//	for i := 0; i < 5; i++ {
//		i := i
//		go func() {
//			defer wg.Done()
//			m.Lock()
//			b := bank_network_call(i)
//
//			x[i] = b
//			m.Unlock()
//		}()
//	}
//
//	wg.Wait()
//
//	// Как-то считается сумма значений в мапе и возвращается
//	return sumOfMap
//}
