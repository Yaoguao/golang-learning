package main

import (
	"fmt"
	"time"
)

//	Задание 2
//	Написать программу, которая конкурентно рассчитает значение квадратов чисел
//	взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	arr := []int{2, 4, 6, 8, 10}

	for _, val := range arr {
		// Оборачиваем вызов анонимной функции в отдельный поток(горутину)
		go func(n int) {
			res := n * n
			fmt.Println(res)
		}(val)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Program has completed its work")
}
