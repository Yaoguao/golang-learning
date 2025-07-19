package _map

func WordCount(s []string) map[string]int {
	res := make(map[string]int)
	for _, val := range s {
		if i, ok := res[val]; ok {
			res[val] = i + 1
			continue
		}
		res[val] = 1
	}
	return res
}

func InvertMap(m map[string]string) map[string][]string {
	res := make(map[string][]string)
	for key, val := range m {
		i, _ := res[val]
		res[val] = append(i, key)
	}
	return res
}

func MapsEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, val := range a {
		if i, ok := b[key]; !ok || i != val {
			return false
		}
	}
	return true
}

/*
USING reflect.DeepEqual(a, b)
*/
//func MapsEqual(a, b map[string]int) bool {
//	return reflect.DeepEqual(a, b)
//}
