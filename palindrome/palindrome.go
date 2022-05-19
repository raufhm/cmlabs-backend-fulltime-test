package palindrome

import (
	"fmt"
	"strings"
)

func Solution(inputString string) bool {
	var penampung string
	inp := strings.ToLower(strings.Replace(inputString, " ", "", -1))
	input := []rune(inp)
	for i := range input {
		penampung = fmt.Sprintf("%s%s", penampung, string(input[len(input)-i-1]))
	}
	if penampung == inp {
		return true
	}
	return false
}
