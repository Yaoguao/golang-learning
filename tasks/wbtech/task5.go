package main

import "fmt"

//	Задание 5
//	Разработать программу, которая будет последовательно отправлять значения в
//	канал, а с другой стороны канала — читать. По истечению N секунд программа
//	должна завершаться.

func pushing(n int, ch chan int) {
	ch <- n
}

func put(ch chan int) int {
	return <-ch
}

func main() {
	var second int
	fmt.Println("Введите через сколько(сек) программа завершиться: ")
	_, err := fmt.Scanln(&second)
	if err != nil {
		return
	}

}
