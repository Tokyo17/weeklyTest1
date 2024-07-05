package quiz

import "sort"

func SliceOperation(nums1 []int, nums2 []int, typeOperation int) []int {
	result := []int{}
	if typeOperation == 1 {
		result = type1(nums1, nums2)
	} else if typeOperation == 2 {
		result = type2(nums1, nums2)
	} else if typeOperation == 3 {
		result = type3(nums1, nums2)
	} else if typeOperation == 4 {
		result = type4(nums1, nums2)
	}
	return result
}

func type1(nums1 []int, nums2 []int) []int {
	result := []int{}
	for _, num1 := range nums1 {
		same := false
		for _, num2 := range nums2 {
			if num1 == num2 {
				same = true
			}
		}
		if !same {
			result = append(result, num1)
		}
	}
	return result
}

func type2(nums1 []int, nums2 []int) []int {
	result := []int{}
	for _, num2 := range nums2 {
		same := false
		for _, num1 := range nums1 {
			if num1 == num2 {
				same = true
			}
		}
		if !same {
			result = append(result, num2)
		}
	}
	return result
}

func type3(nums1 []int, nums2 []int) []int {
	filter := map[int]bool{}
	result := []int{}
	newArr := []int{}
	newArr = append(newArr, nums1...)
	newArr = append(newArr, nums2...)

	for _, v := range newArr {
		filter[v] = true
	}
	for i, _ := range filter {
		result = append(result, i)
	}
	sort.Ints(result)
	return result

}

func type4(nums1 []int, nums2 []int) []int {
	result := []int{}
	for _, num2 := range nums2 {
		same := false
		for _, num1 := range nums1 {
			if num1 == num2 {
				same = true
			}
		}
		if same {
			result = append(result, num2)
		}
	}
	return result
}
