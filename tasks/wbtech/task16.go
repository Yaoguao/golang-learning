package main

import "fmt"

// Задание 16
// Реализовать быструю сортировку массива (quicksort) встроенными методами
// языка.
// Можно ту и ту функцию использовать так-то
func QSort(sl []int) []int {
	res := make([]int, len(sl))
	copy(res, sl)
	fmt.Println(sl)
	quicksort(res, 0, len(res)-1)
	return res
}

func quicksort(sl []int, low, high int) {
	if low < high {
		p := func(sl []int, low, high int) int { // можно вывести в отдельную функцию, хз
			pivot := sl[(low+high)/2]
			i := low
			j := high
			for {
				for sl[i] < pivot {
					i++
				}
				for sl[j] > pivot {
					j--
				}
				if i >= j {
					return j
				}
				sl[i], sl[j] = sl[j], sl[i]
				i++
				j--
			}
		}(sl, low, high)
		quicksort(sl, low, p)
		quicksort(sl, p+1, high)
	}
}

func main() {
	sl := []int{2, 5, 7, 9, 4, 10, 43, 5, 1, 788, 56, 43, 12, 8, 767, 4, 6, 7, 9, 5}
	sl = QSort(sl)
	fmt.Println(sl)
}
