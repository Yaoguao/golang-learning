package main

import (
	"fmt"
)

// Задача 14
// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

func determinant(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan any: // ????
		fmt.Println("chan")
	default:
		fmt.Println("I dont know")
	}
}

// Вариант с рефлексией, тогда можно определить любой канал
//func determinant(i interface{}) {
//	switch v := reflect.ValueOf(i); v.Kind() {
//	case reflect.Int:
//		fmt.Println("int")
//	case reflect.String:
//		fmt.Println("string")
//	case reflect.Bool:
//		fmt.Println("bool")
//	case reflect.Chan:
//		fmt.Println("chan")
//	default:
//		fmt.Println("I dont know")
//	}
//}

func main() {
	//ch := make(chan string)
	determinant(45)
}
