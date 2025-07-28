package main

import "fmt"

//	Реализовать пересечение двух неупорядоченных множеств.

func Intersection(nums1 []int, nums2 []int) []int {
	var numsMap = make(map[int]int)
	var res []int

	for _, num := range nums1 {
		if numsMap[num] > 0 {
			continue
		} else {
			numsMap[num] = 1
		}
	}

	for _, num := range nums2 {
		if numsMap[num] > 0 {
			numsMap[num] += 1
		} else {
			continue
		}
	}

	for num, count := range numsMap {
		if count > 1 {
			res = append(res, num)
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 6, 6, 3, 5, 8, 9, 9}

	nums2 := []int{3, 8, 2, 9, 1, 9}

	fmt.Println(Intersection(nums1, nums2))
}
