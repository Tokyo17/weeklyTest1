package quiz

func OddBeforeEven(nums []int) []int {
	result := []int{}
	odd := []int{}
	even := []int{}
	for _, v := range nums {
		if v%2 != 0 {
			odd = append(odd, v)
		} else {
			even = append(even, v)
		}
	}
	result = append(result, odd...)
	result = append(result, even...)
	return result
}
