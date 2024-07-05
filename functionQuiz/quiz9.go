package quiz

func CountFrequentElement(nums []int) map[int]int {
	result := map[int]int{}
	for _, v := range nums {
		result[v]++
	}
	return result
}
