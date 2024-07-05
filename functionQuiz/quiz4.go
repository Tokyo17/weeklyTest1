package quiz

func RemoveDuplicate(num []int) []int {
	newMap := map[int]bool{}
	result := []int{}
	for _, v := range num {
		newMap[v] = true
	}
	for i, _ := range newMap {
		result = append(result, i)
	}
	return result
}
