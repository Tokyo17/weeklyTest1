package quiz

func SliceFruits(fruits1 []string, fruits2 []string, operationType int) []string {
	result := []string{}
	if operationType == 1 {
		result = opType1(fruits1, fruits2)
	} else if operationType == 2 {
		result = opType2(fruits1, fruits2)
	}
	return result
}

func opType1(fruits1 []string, fruits2 []string) []string {
	result := []string{}
	filter := map[string]bool{}
	for _, fruit1 := range fruits1 {
		same := false
		for _, fruit2 := range fruits2 {
			if fruit1 == fruit2 {
				same = true
			}
		}
		if same {
			filter[fruit1] = true
		}

	}

	for i, _ := range filter {
		result = append(result, i)
	}
	return result

}

func opType2(fruits1 []string, fruits2 []string) []string {
	result := []string{}
	for _, fruit1 := range fruits1 {
		same := false
		for _, fruit2 := range fruits2 {
			if fruit1 == fruit2 {
				same = true
			}
		}
		if !same {
			result = append(result, fruit1)
		}

	}

	for _, fruit2 := range fruits2 {
		same := false
		for _, fruit1 := range fruits1 {
			if fruit1 == fruit2 {
				same = true
			}
		}
		if !same {
			result = append(result, fruit2)
		}

	}

	return result

}
