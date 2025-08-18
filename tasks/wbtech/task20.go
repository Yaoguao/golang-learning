package main

import (
	"fmt"
	"strings"
)

// Задание 20
//
//	Разработать программу, которая переворачивает слова в строке.
//	Пример: «snow dog sun — sun dog snow».
//

// Решение как с прошлой задачей
func ReverseWords(str string) string {
	// Либо strings.Split()
	sl := strings.Fields(str)
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return strings.Join(sl, " ")
}

// Решение с билдером
//func ReverseWords(str string) string {
//	var builder strings.Builder
//	// Либо strings.Split()
//	sl := strings.Fields(str)
//	for i := len(sl) - 1; i >= 0; i-- {
//		builder.WriteString(sl[i])
//		builder.WriteString(" ")
//	}
//	return strings.TrimSpace(builder.String())
//}

func main() {
	str := "snow dog sun"
	fmt.Println(ReverseWords(str))
}
