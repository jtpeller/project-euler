// ============================================================================
// = problems.go
// = 	Description: implementation for the project euler problems
// = 	Problem:
// = 	Date: December 16, 2021
// ============================================================================

package prob

import (
	"fmt"
	"math"
	"project-euler/utils"
)

/*
If we list all the natural numbers below 10 that are multiples of
3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 Find the sum of all the multiples of 3 or 5 below 1000.
*/
func P1() int64 {
	m := make([]int64, 0)
	for i := int64(0); i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			m = append(m, i)
		}
	}
	fmt.Println(m)
	return utils.Sum(m)
}

/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
 By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/
func P2() int64 {
	F := utils.Fibonacci(35)
	sum := int64(0)
	for _, f := range F {
		if f % 2 == 0 && f < 4e6 {
			sum += f
		}
	}
	return sum
}

/*
The prime factors of 13195 are 5, 7, 13 and 29.
 What is the largest prime factor of the number 600851475143 ?
*/
func P3() int64 {
	fact := utils.PrimeFactorization(600851475143)
	fmt.Println(fact)
	return fact[len(fact)-1]
}

/*
A palindromic number reads the same both ways. The largest 
palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
 Find the largest palindrome made from the product of two 3-digit numbers.
*/
func P4() int64 {
	min := int64(100)
	max := int64(999)
	mp := int64(0)
	n1 := int64(0)
	n2 := int64(0)
	prod := int64(0)
	for i := max; i >= min; i-- {
		for j := max; j >= min; j-- {
			prod = int64(i * j)
			if utils.IsPalindrome(prod) && mp < prod {
				n1 = i
				n2 = j
				mp = prod
			}
		}
	}
	fmt.Println("factors:", n1, ",", n2)
	return mp
}

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
 What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
func P5() int64 {
	max := int64(20)
	num := int64(20)		// start with 20
	found := false
	for !found {
		for i := int64(1); i < max; i++ {
			if num % i != 0 {
				found = false
				num++
				break;
			}
			found = true
		}
	}
	return num
}

/*
The sum of the squares of the first ten natural numbers is,
 1^2 + 2^2 +...+10^2 = 385
The square of the sum of the first ten natural numbers is,
 (1+2+...+10)^2 = 55^2 = 3025
Hence the difference between the sum of the squares of the first
ten natural numbers and the square of the sum is
 3025-385=2640

 Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
func P6() int64 {
	n := int64(100)
	nf := float64(n)
	return int64(((math.Pow(nf, 2) * math.Pow(nf + 1.0, 2)) / 4) -		// square of the sum
		float64(n * (n + 1) * (2*n + 1) / 6))	// sum of squares
}

/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
 What is the 10 001st prime number?
*/
func P7() int64 {
	primes := utils.Primes(10001)
	return primes[10000]
}

/*
The four adjacent digits in the 1000-digit number that have the 
greatest product are 9 × 9 × 8 × 9 = 5832.
 Find the thirteen adjacent digits in the 1000-digit number 
 that have the greatest product. What is the value of this product?
*/
func P8() int64 {
	p8 := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"
	p8sep := utils.Digits(p8)
	mp := int64(0)
	ndigits := 13
	var which []int64
	for i := 0; i < len(p8sep)-ndigits; i++ {
		prod := utils.Prod(p8sep[i:i+ndigits])
		if prod > mp {
			mp = prod
			which = p8sep[i:i+ndigits]
		}
	}
	fmt.Println(which)
	return mp
}

/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.
There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 Find the product abc.
*/
func P9() int64 {
	triplet := func(sum int64) (int64, int64, int64) {
		c := int64(0)
		for a := int64(1); a < sum; a++ {
			af := float64(a)
			for b := int64(1); b < sum; b++ {
				bf := float64(b)
				foo := math.Sqrt(math.Pow(af, 2) + math.Pow(bf, 2))
				if foo == math.Floor(foo) {
					c = int64(foo)
					if a + b + c == sum {
						return a, b, c
					}
				}
			}
		}
		return -1, -1, -1
	}
	a, b, c := triplet(1000)
	return a*b*c
}

/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 Find the sum of all the primes below two million.
*/
func P10() int64 {
	return utils.Sum(utils.PrimesUpTo(2e6))
}

/*
In the 20×20 grid below, four numbers along a diagonal line have been marked in red.
The product of these numbers is 26 × 63 × 78 × 14 = 1788696.
 What is the greatest product of four adjacent numbers in the same direction 
 (up, down, left, right, or diagonally) in the 20×20 grid?
*/
func P11() int64 {
	return 0
}