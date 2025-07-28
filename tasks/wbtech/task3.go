package main

import (
	"fmt"
	"sync"
)

//	Задание 3
//	Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
//	квадратов(22+32+42….) с использованием конкурентных вычислений.

// V1
func SqNumber(n int) int {
	return n * n
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var res int
	var wg sync.WaitGroup

	for _, val := range arr {
		wg.Add(1)
		go func(n int) {
			res += SqNumber(n)
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println(res)
}

// V2 Hmm
//func SqNumber(n int, ch chan int) {
//	ch <- n * n
//}
//
//func main() {
//	arr := []int{2, 4, 6, 8, 10}
//	var res int
//	ch := make(chan int)
//
//	for _, val := range arr {
//		go SqNumber(val, ch)
//		res += <-ch
//	}
//	fmt.Println(res)
//}
