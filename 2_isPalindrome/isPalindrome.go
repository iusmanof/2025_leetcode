package ispalindrome

import (
	"strconv"
)

func ispalindrome(x int) bool {
	if x < 0 {
		return false
	}

	stringX := strconv.Itoa(x)

	for i, j := 0, len(stringX)-1; i < j; i, j = j+1, j-1 {
		if stringX[i] != stringX[j] {
			return false
		}
	}

	return true
}
