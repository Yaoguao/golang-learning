package main

import "fmt"

//	Задание 17
//	Реализовать бинарный поиск встроенными методами языка.
//
// return index
func BinarySearch(val int, sl []int) int {
	left, right, mid := 0, len(sl)-1, 0
	for left <= right {
		mid = left + (right-left)/2
		if sl[mid] == val {
			return mid
		} else if sl[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BinarySearch(3, sl))
}
