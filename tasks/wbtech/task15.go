package main

import "fmt"

//	Задание 15
//	К каким негативным последствиям может привести данный фрагмент кода, и как
//	это исправить? Приведите корректный пример реализации.
//
//	var justString string
//	func someFunc() {
//		v := createHugeString(1 << 10)
//		justString = v[:100]
//	}
//	func main() {
//		someFunc()
//	}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// Подставить копию? Типо данные ссылются на одну область пямяти
	justString = v[:100]

	fmt.Println(justString)
}

func createHugeString(len int) string {
	return string(make([]byte, len))
}

func main() {
	someFunc()
}
