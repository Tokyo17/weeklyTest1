package quiz

import (
	"strings"
)

func ExceptCapitalize(words []string, except []string) []string {
	result := []string{}
	for _, word := range words {
		cap := true
		for _, exc := range except {
			if word == exc {
				cap = false
			}
		}
		if cap {
			result = append(result, strings.ToUpper(string(word[0]))+word[1:])
		} else {
			result = append(result, word)
		}
	}
	return result
}
