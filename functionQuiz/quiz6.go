package quiz

func PlusOne(nums []int) []int {
	resultSum := []int{}
	sisa := false
	for i := len(nums) - 1; i >= 0; i-- {
		if !sisa {
			if (nums[i]+1)/10 != 0 {
				resultSum = append(resultSum, 0)
				sisa = true
				if i == 0 {
					resultSum = append(resultSum, 1)
				}
			} else {
				resultSum = append(resultSum, nums[i])
			}
		} else {
			if (nums[i]+1)/10 != 0 {
				resultSum = append(resultSum, 0)
				sisa = true
				if i == 0 {
					resultSum = append(resultSum, 1)
				}
			} else {
				resultSum = append(resultSum, nums[i]+1)
				sisa = false
			}
		}

	}

	finalResult := []int{}

	for i := len(resultSum) - 1; i >= 0; i-- {
		finalResult = append(finalResult, resultSum[i])
	}

	return finalResult
}
