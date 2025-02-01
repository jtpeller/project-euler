// ============================================================================
// = calc.go
// = 	Description		various calculation functions (say, multiples)
// = 	Date			December 16, 2021
// ============================================================================

package utils

import (
	"math"
	"math/big"
	"strconv"

	gb "github.com/jtpeller/gobig"
)

// computs C(n, k)
func C(n, k int64) int64 {
	if n < 0 || k < 0 {
		PrintError("cannot be negative")
	} else if n < k {
		PrintError("n cannot be less than k")
	}

	// shortcut for speed
	if k > n/2 {
		k = n - k
	}

	c := int64(1)
	for i := int64(1); i <= k; i++ {
		c = (n - k + i) * c / i
	}
	return c
}

// separates a number into its digits
func Digits(num string) []int64 {
	ndigits := len(num)
	a := make([]int64, ndigits)
	for i := 0; i < ndigits; i++ {
		d, _ := strconv.ParseInt(string(num[i]), 10, 64)
		a[i] = d
	}
	return a
}

// computes the divisors of the provided number
func Divisors(num int64) []int64 {
	div := make([]int64, 0)
	sqrt_num := int64(math.Sqrt(float64(num)))
	for i := int64(1); i <= sqrt_num; i++ {
		if num % i == 0 {
			// if this is a sqrt, only append i
			if (num / i == i) {
				div = append(div, i)
			} else {
				div = append(div, i, num/i)
			}
		}
	}
	return div
}

func Factorial(num int64) int64 {
	prod := int64(1)
	for i := int64(2); i <= num; i++ {
		prod *= i
	}
	return prod
}

// computes the fibonacci numbers
func Fibonacci(count int64) []int64 {
	a := make([]int64, count)
	a[0] = 0
	a[1] = 1
	for i := int64(2); i < count; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	return a
}

// computes max of array
func Max(a []int64) (int64, int) {
	max := int64(-1 << 63)
	maxidx := 0
	for idx, num := range a {
		if (max < num) {
			max = num
			maxidx = idx
		}
	}
	return max, maxidx
}

// Calculates the prime factorization of num
func PrimeFactorization(num int64) []int64 {
	// initialize array
	primefact := make([]int64, 0)

	// the number of 2s
	for num % 2 == 0 {
		primefact = append(primefact, 2)
		num /= 2
	}

	// num is now odd. check 3s and beyond
	for i := int64(3); i * i <= num; i += 2 {
		// get factors until n is zero
		for num % i == 0 {
			primefact = append(primefact, i)
			num /= i
		}
	}

	// what if num is prime?
	if num > 2 {
		primefact = append(primefact, num)
	}

	return primefact
}

// computes the n primes
func Primes(n int64) []int64 {
	primes := make([]int64, 0)
	num := int64(0)
	for i := int64(0); i < n; {
		if IsPrime(num) {
			primes = append(primes, num)
			i++
		}
		num++
	}
	return primes
}

// computes primes up to max
func PrimesUpTo(max int64) []int64 {
	primes := make([]int64, 0)
	for i := int64(1); i < max; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// computes the product of the terms in the array, like Sum(), but for multiplication
func Prod(a []int64) int64 {
	prod := int64(1)
	for i := 0; i < len(a); i++ {
		prod *= a[i]
	}
	return prod
}

// calculates the sum of a given array
func Sum(a []int64) int64 {
	sum := int64(0)
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

// computes Sum(), except with *big.Int
func SumBig(a []*big.Int) *big.Int {
	sum := gb.New(0)
	for i := 0; i < len(a); i++ {
		sum.Add(sum, a[i])
	}
	return sum
}

// calculates triangle number
func Triangle(a int64) int64 {
	// there is a formula to compute in O(1) time
	return a*(a+1)/2
}