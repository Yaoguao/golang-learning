package main

import "fmt"

//	Задание 13
//	Поменять местами два числа без создания временной переменной.

func main() {
	a, b := 13, 15
	fmt.Printf("a: %d, b: %d\n", a, b) // a: 13, b: 15
	a, b = b, a
	fmt.Printf("a: %d, b: %d", a, b) // a: 15, b: 13

}
