package quiz

import (
	"reflect"
	"sort"
	"strings"
)

func IsAnagram(input1, input2 string) bool {
	a := []string{}
	b := []string{}

	for i := 0; i < len(input1); i++ {
		a = append(a, strings.ToLower(string(input1[i])))
		b = append(b, strings.ToLower(string(input2[i])))
	}
	sort.Strings(a)
	sort.Strings(b)
	return reflect.DeepEqual(a, b)

}
