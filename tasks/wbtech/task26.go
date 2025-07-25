package main

import (
	"fmt"
	"strings"
)

//	Задание 26
//	Разработать программу, которая проверяет, что все символы в строке
//	уникальные (true — если уникальные, false etc). Функция проверки должна быть
//	регистронезависимой.
//	Например:
//	abcd — true
//	abCdefAaf — false
//	aabcd — false

func Unique(s string) bool {
	ch := make(map[rune]struct{})
	for _, val := range strings.ToLower(s) {
		if _, ok := ch[val]; ok {
			return false
		}
		ch[val] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(Unique("acfgA"))
}
