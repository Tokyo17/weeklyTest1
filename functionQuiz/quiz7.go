package quiz

func SameElement(nums1 []int, nums2 []int) []int {

	result := []int{}

	for _, num1 := range nums1 {
		duplicate := false
		for _, num2 := range nums2 {
			if num1 == num2 {
				duplicate = true
			}
		}
		if duplicate {
			result = append(result, num1)
		}
	}

	return result
}
