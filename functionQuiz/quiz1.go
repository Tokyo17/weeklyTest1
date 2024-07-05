package quiz

import (
	"strings"
)

func RemoveDuplicateLetter(input string) string {
	filter := map[string]bool{}
	result := []string{}
	for i := 0; i < len(input); i++ {
		filter[string(input[i])] = true
	}
	for i, _ := range filter {
		result = append(result, i)
	}
	// fmt.Println()
	return strings.Join(result, "")
}
