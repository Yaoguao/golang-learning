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

// V1
func Unique(s []string) []string {
	res := make([]string, 0, len(s))
	ch := make(map[string]string)
	for _, val := range s {
		if _, ok := ch[val]; ok {
			continue
		}
		ch[val] = val
		res = append(res, val)
	}
	return res
}

// V2 O(n^2) :D
//func Unique(s []string) []string {
//	res := make([]string, 0, len(s))
//	isUniq := false
//	for i := 0; i < len(s); i++ {
//		for j := i + 1; j < len(s); j++ {
//			if s[i] == s[j] {
//				isUniq = true
//				break
//			}
//		}
//		if isUniq {
//			isUniq = false
//			continue
//		}
//		res = append(res, s[i])
//	}
//	return res
//}

func InsertAt(s []int, index int, value int) []int {
	res := make([]int, 0, len(s)+1)
	for idx, elem := range s {
		if index == idx || index < 0 {
			res = append(res, value)
		}
		res = append(res, elem)
	}
	if index >= len(s) {
		res = append(res, value)
	}
	return res
}
