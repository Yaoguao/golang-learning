package main

import (
	"fmt"
	"sync"
)

//	Задание 2
//	Написать программу, которая конкурентно рассчитает значение квадратов чисел
//	взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, val := range arr {
		wg.Add(1)
		// Оборачиваем вызов анонимной функции в отдельный поток(горутину)
		go func(n int) {
			fmt.Println(n * n)
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println("Program has completed its work")
}
