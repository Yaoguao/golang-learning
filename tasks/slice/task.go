package slice

// V1
func RemoveAtIndex(s []int, index int) []int {
	res := make([]int, 0, len(s))
	for idx, val := range s {
		if idx == index {
			continue
		}
		res = append(res, val)
	}
	return res
}
