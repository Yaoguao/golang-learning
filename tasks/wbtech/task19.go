package main

import "fmt"

//	Задание 19
//	Разработать программу, которая переворачивает подаваемую на ход строку
//	(например: «главрыба — абырвалг»). Символы могут быть unicode.

func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// V2 чисто на байтах, пусть будет
//func Reverse(str string) string {
//	var ridx int
//	res := make([]byte, len(str))
//	for i := len(str) - 1; i >= 0; i-- {
//		res[ridx] = str[i]
//		ridx++
//	}
//	return string(res)
//}

func main() {
	fmt.Println(Reverse("макс режнессеМ")) // Мессенжер скам
}
