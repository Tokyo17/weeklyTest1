package quiz

import "fmt"

func AddDigits(nums1 []int, nums2 []int) []int {

	nextNum := 0
	resultSum := []int{}

	for i := len(nums1) - 1; i >= 0; i-- {
		if nextNum == 0 {
			if (nums1[i]+nums2[i])/10 != 0 {
				nextNum = (nums1[i] + nums2[i]) / 10
				resultSum = append(resultSum, (nums1[i]+nums2[i])%10)
				if i == 0 {
					resultSum = append(resultSum, nextNum)
				}
			} else {
				resultSum = append(resultSum, nums1[i]+nums2[i])
			}
		} else {
			if (nums1[i]+nums2[i]+nextNum)/10 != 0 {
				fmt.Println(nums1[i], nums2[i])
				nextNum = (nums1[i] + nums2[i] + nextNum) / 10
				resultSum = append(resultSum, (nums1[i]+nums2[i])%10)
				if i == 0 {
					resultSum = append(resultSum, nextNum)
				}
			} else {
				resultSum = append(resultSum, nums1[i]+nums2[i]+nextNum)
				nextNum = 0
			}
		}
	}
	finalResult := []int{}

	for i := len(resultSum) - 1; i >= 0; i-- {
		finalResult = append(finalResult, resultSum[i])
	}
	return finalResult
}
