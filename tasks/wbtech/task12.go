package main

import "fmt"

// Задание 12
// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.
// *Тут можно map[string]string, что бы удобно было читать, но я хз
func newSet(sl []string) map[string]struct{} {
	res := make(map[string]struct{}, len(sl))
	for _, val := range sl {
		res[val] = struct{}{}
	}
	return res
}

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	set := newSet(sl)
	for key, _ := range set { // только так читаем
		fmt.Println(key)
	}
	fmt.Println(set)
}
