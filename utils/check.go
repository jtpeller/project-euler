// ============================================================================
// = check.go
// = 	Description: various checker functions (say, IsPrime)
// = 	Date: December 16, 2021
// ============================================================================

package utils

import (
	"math/big"
	"strconv"
)

func IsPalindrome(n int64) bool {
	str := strconv.FormatInt(n, 10)
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

// IsPrime returns true if num is prime. False otherwise.
func IsPrime(num int64) bool {
	return big.NewInt(num).ProbablyPrime(20)
}
