// ============================================================================
// = problems.go
// = 	Description		Implementation for the project euler problems
// = 	Date			2021.12.16
// ============================================================================

package prob

import (
	"fmt"
	"math"
	"project-euler/utils"
	"strconv"
	"strings"
	"time"

	n2w "github.com/jtpeller/num2words"
)

/**
 * P1() - Multiples of 3 or 5
 *  Find the sum of all the multiples of 3 or 5 below 1000.
 *	Date	2021.01.20
 */
func P1() int64 {
	m := make([]int64, 0)
	for i := int64(0); i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			m = append(m, i)
		}
	}
	fmt.Println(m)
	return utils.Sum(m)
}

/**
 * P2() - Even Fibonacci Numbers
 *  By considering the terms in the Fibonacci sequence whose values
 * 		do not exceed four million, find the sum of the even-valued
 *		terms.
 *	Date	2021.01.20
 */
func P2() int64 {
	F := utils.Fibonacci(35)
	sum := int64(0)
	for _, f := range F {
		if f%2 == 0 && f < 4e6 {
			sum += f
		}
	}
	return sum
}

/**
 * P3() - Largest Prime Factor
 *  What is the largest prime factor of the number 600851475143?
 *	Date	2021.01.21
 */
func P3() int64 {
	fact := utils.PrimeFactorization(600851475143)
	fmt.Println(fact)
	return fact[len(fact)-1]
}

/**
 * P4() - Largest Palindrome Product
 *  Find the largest palindrome made from the product of two 3-digit
 *		numbers.
 *	Date	2021.01.21
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

/**
 * P5() - Smallest Multiple
 *  What is the smallest positive number that is evenly divisible
 *		by all of the numbers from 1 to 20?
 *	Date	2021.01.21
 */
func P5() int64 {
	max := int64(20)
	num := int64(20) // start with 20
	found := false
	for !found {
		for i := int64(1); i < max; i++ {
			if num%i != 0 {
				found = false
				num++
				break
			}
			found = true
		}
	}
	return num
}

/**
 * P6() - Sum Square Difference
 *  Find the difference between the sum of the squares of the first
 * 		one hundred natural numbers and the square of the sum.
 *	Date	2021.01.23
 */
func P6() int64 {
	n := int64(100)
	nf := float64(n)
	return int64(((math.Pow(nf, 2) * math.Pow(nf+1.0, 2)) / 4) - // square of the sum
		float64(n*(n+1)*(2*n+1)/6)) // sum of squares
}

/**
 * P7() - 10,001st Prime
 * What is the 10,001st prime number?
 *	Date	2021.01.23
 */
func P7() int64 {
	primes := utils.Primes(10001)
	return primes[10000]
}

/**
 * P8() - Largest Product in a Series
 *  Find the thirteen adjacent digits in the 1000-digit number that
 *		have the greatest product. What is the value of this product?
 *	Date	2021.01.24
 */
func P8() int64 {
	p8 := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"
	p8sep := utils.Digits(p8)
	mp := int64(0)
	ndigits := 13
	var which []int64
	for i := 0; i < len(p8sep)-ndigits; i++ {
		prod := utils.Prod(p8sep[i : i+ndigits])
		if prod > mp {
			mp = prod
			which = p8sep[i : i+ndigits]
		}
	}
	fmt.Println(which)
	return mp
}

/**
 * P9() - Special Pythagorean Triplet
 *  There exists exactly one Pythagorean triplet for which
 *		a + b + c = 1000. Find the product abc.
 *	Date	2021.01.24
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
					if a+b+c == sum {
						return a, b, c
					}
				}
			}
		}
		return -1, -1, -1
	}
	a, b, c := triplet(1000)
	return a * b * c
}

/**
 * P10() - Summation of Primes
 *  Find the sum of all the primes below two million.
 *	Date	2021.01.24
 */
func P10() int64 {
	return utils.Sum(utils.PrimesUpTo(2e6))
}

/**
 * P11() - Largest Product in a Grid
 * 	What is the greatest product of four adjacent numbers in the
 *		same direction (up, down, left, right, or diagonally) in the
 *		20×20 grid?
 *	Date	2025.01.28
 */
func P11() int64 {
	// raw is pulled from the site, and modified for storage
	raw := "08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08\n49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00\n81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65\n52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91\n22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80\n24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50\n32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70\n67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21\n24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72\n21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95\n78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92\n16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57\n86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58\n19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40\n04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66\n88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69\n04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36\n20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16\n20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54\n01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48"
	count := 4 // how many consecutive #s to compute the product for

	// first, separate by newline
	rows := strings.Split(raw, "\n")
	nrows := len(rows)
	ncols := nrows // square grid

	// then loop through rows

	// ... make the grid
	grid := make([][]int64, nrows)
	for i := range grid {
		grid[i] = make([]int64, nrows)
	}

	// ... extract values
	for i, row := range rows {
		nums := strings.Split(row, " ")
		for j, num := range nums {
			grid[i][j], _ = strconv.ParseInt(num, 10, 64)
		}
	}

	// process grid to compute maximum product
	product := int64(0)
	factors := make([]int64, count)

	// ... left / right
	for row := 0; row < nrows; row++ {
		for col := 0; col < ncols-count; col++ {
			vals := grid[row][col : col+count]
			prod := utils.Prod(vals)

			// check if prod is bigger than max
			if prod > product {
				product = prod
				copy(factors, vals)
			}
		}
	}
	fmt.Println("Highest value from left/right", product, factors)

	// ... up / down
	for col := 0; col < ncols; col++ {
		for row := 0; row < nrows-count; row++ {
			// extract vertical nums
			vals := make([]int64, count)
			for i := 0; i < count; i++ {
				vals[i] = grid[row+i][col]
			}

			prod := utils.Prod(vals)

			// check if prod is bigger than max
			if prod > product {
				product = prod
				copy(factors, vals)
			}
		}
	}
	fmt.Println("Highest value from up/down", product, factors)

	// ... diagonals (top left to bottom right)
	for row := 0; row < nrows-count; row++ {
		for col := 0; col < ncols-count; col++ {
			// extract vertical nums
			vals := make([]int64, count)
			for i := 0; i < count; i++ {
				vals[i] = grid[row+i][col+i]
			}
			prod := utils.Prod(vals)

			// check if prod is bigger than max
			if prod > product {
				product = prod
				copy(factors, vals)
			}
		}
	}
	fmt.Println("Highest value from diag top left -> bottom right", product, factors)

	// ... diagonals (bottom left to top right)
	for row := nrows - 1; row >= count; row-- {
		for col := 0; col < ncols-count; col++ {
			// extract vertical nums
			vals := make([]int64, count)
			for i := 0; i < count; i++ {
				vals[i] = grid[row-i][col+i]
			}
			prod := utils.Prod(vals)

			// check if prod is bigger than max
			if prod > product {
				product = prod
				copy(factors, vals)
			}
		}
	}

	fmt.Println("Product and factors after computation", product, factors)
	return product
}

/**
 * P12() - Highly Divisible Triangular Number.
 *  What is the value of the first triangle number to have over five
 *		hundred divisors?
 *	Date	2025.01.28
 */
func P12() int64 {
	// init
	tnum := int64(0)
	var div []int64

	// begin the loop
	idx := int64(0)
	for {
		tnum = utils.Triangle(idx)
		div = utils.Divisors(tnum)
		if len(div) >= 500 {
			break
		}
		idx++
	}

	// print
	fmt.Println("Num:", tnum)
	fmt.Println("Index:", idx)
	//fmt.Println("Divisors:", div)		// this prints too much

	// return
	return tnum
}

/**
 *	P13() - Large Sum
 * 	 Work out the first ten digits of the sum of the following
 *		one-hundred 50-digit numbers.
 *	Date	2025.01.28
 */
func P13() int64 {
	// take values
	raw := "37107287533902102798797998220837590246510135740250\n46376937677490009712648124896970078050417018260538\n74324986199524741059474233309513058123726617309629\n91942213363574161572522430563301811072406154908250\n23067588207539346171171980310421047513778063246676\n89261670696623633820136378418383684178734361726757\n28112879812849979408065481931592621691275889832738\n44274228917432520321923589422876796487670272189318\n47451445736001306439091167216856844588711603153276\n70386486105843025439939619828917593665686757934951\n62176457141856560629502157223196586755079324193331\n64906352462741904929101432445813822663347944758178\n92575867718337217661963751590579239728245598838407\n58203565325359399008402633568948830189458628227828\n80181199384826282014278194139940567587151170094390\n35398664372827112653829987240784473053190104293586\n86515506006295864861532075273371959191420517255829\n71693888707715466499115593487603532921714970056938\n54370070576826684624621495650076471787294438377604\n53282654108756828443191190634694037855217779295145\n36123272525000296071075082563815656710885258350721\n45876576172410976447339110607218265236877223636045\n17423706905851860660448207621209813287860733969412\n81142660418086830619328460811191061556940512689692\n51934325451728388641918047049293215058642563049483\n62467221648435076201727918039944693004732956340691\n15732444386908125794514089057706229429197107928209\n55037687525678773091862540744969844508330393682126\n18336384825330154686196124348767681297534375946515\n80386287592878490201521685554828717201219257766954\n78182833757993103614740356856449095527097864797581\n16726320100436897842553539920931837441497806860984\n48403098129077791799088218795327364475675590848030\n87086987551392711854517078544161852424320693150332\n59959406895756536782107074926966537676326235447210\n69793950679652694742597709739166693763042633987085\n41052684708299085211399427365734116182760315001271\n65378607361501080857009149939512557028198746004375\n35829035317434717326932123578154982629742552737307\n94953759765105305946966067683156574377167401875275\n88902802571733229619176668713819931811048770190271\n25267680276078003013678680992525463401061632866526\n36270218540497705585629946580636237993140746255962\n24074486908231174977792365466257246923322810917141\n91430288197103288597806669760892938638285025333403\n34413065578016127815921815005561868836468420090470\n23053081172816430487623791969842487255036638784583\n11487696932154902810424020138335124462181441773470\n63783299490636259666498587618221225225512486764533\n67720186971698544312419572409913959008952310058822\n95548255300263520781532296796249481641953868218774\n76085327132285723110424803456124867697064507995236\n37774242535411291684276865538926205024910326572967\n23701913275725675285653248258265463092207058596522\n29798860272258331913126375147341994889534765745501\n18495701454879288984856827726077713721403798879715\n38298203783031473527721580348144513491373226651381\n34829543829199918180278916522431027392251122869539\n40957953066405232632538044100059654939159879593635\n29746152185502371307642255121183693803580388584903\n41698116222072977186158236678424689157993532961922\n62467957194401269043877107275048102390895523597457\n23189706772547915061505504953922979530901129967519\n86188088225875314529584099251203829009407770775672\n11306739708304724483816533873502340845647058077308\n82959174767140363198008187129011875491310547126581\n97623331044818386269515456334926366572897563400500\n42846280183517070527831839425882145521227251250327\n55121603546981200581762165212827652751691296897789\n32238195734329339946437501907836945765883352399886\n75506164965184775180738168837861091527357929701337\n62177842752192623401942399639168044983993173312731\n32924185707147349566916674687634660915035914677504\n99518671430235219628894890102423325116913619626622\n73267460800591547471830798392868535206946944540724\n76841822524674417161514036427982273348055556214818\n97142617910342598647204516893989422179826088076852\n87783646182799346313767754307809363333018982642090\n10848802521674670883215120185883543223812876952786\n71329612474782464538636993009049310363619763878039\n62184073572399794223406235393808339651327408011116\n66627891981488087797941876876144230030984490851411\n60661826293682836764744779239180335110989069790714\n85786944089552990653640447425576083659976645795096\n66024396409905389607120198219976047599490197230297\n64913982680032973156037120041377903785566085089252\n16730939319872750275468906903707539413042652315011\n94809377245048795150954100921645863754710598436791\n78639167021187492431995700641917969777599028300699\n15368713711936614952811305876380278410754449733078\n40789923115535562561142322423255033685442488917353\n44889911501440648020369068063960672322193204149535\n41503128880339536053299340368006977710650566631954\n81234880673210146739058568557934581403627822703280\n82616570773948327592232845941706525094512325230608\n22918802058777319719839450180888072429661980811197\n77158542502016545090413245809786882778948721859617\n72107838435069186155435662884062257473692284509516\n20849603980134001723930671666823555245252804609722\n53503534226472524250874054075591789781264330331690"
	strs := strings.Split(raw, "\n")
	size := int64(len(strs))

	// convert to big int array
	arr := utils.CreateSlice(size)
	for i := int64(0); i < size; i++ {
		arr[i], _ = zero().SetString(strs[i], 10)
	}

	// sum
	sum := utils.SumBig(arr)

	// print sum
	fmt.Println("Sum = ", sum)

	// return just the first 10 digits
	str := sum.String()
	digits, _ := strconv.ParseInt(str[:10], 10, 64)
	return digits
}

/**
 * P14() - Longest Collatz Sequence
 * 	Which starting number, under one million, produces the longest
 *  	chain?
 *	Date	2025.02.01
 */
func P14() int64 {
	// logic is: loop from one million down to zero
	// for each num, compute collatz until finishing at 1
	// save only the length of the collatz chain
	lengths := make([]int64, 0)
	for n := int64(1000000); n >= 1; n-- {
		// compute collatz now
		val := n
		chain := make([]int64, 0)
		for val > 1 {
			if val%2 == 0 {
				val /= 2
			} else {
				val = (3*val + 1)
			}
			chain = append(chain, val)
		}

		// now, save the length
		lengths = append(lengths, int64(len(chain)))
	}

	// now, quickly find the max of length
	max, idx := utils.Max(lengths)
	fmt.Println("Max chain:", max)
	fmt.Println("Number:", 1000000-idx)
	return int64(1000000 - idx)
}

/**
 * P15() - Lattice Paths
 * 	How many routes through a 20x20 grid?
 *	Date	2025.02.01
 */
func P15() int64 {
	return utils.C(40, 20)
}

/**
 * P16() - Power Digit Sum
 *  What is the sum of the digits of 2^1000?
 *	Date	2025.02.01
 */
func P16() int64 {
	num := pow(inew(2), inew(1000)) // compute 2^1000
	fmt.Println("2^1000 =", num)
	digits := utils.Digits(num.String()) // separate into digits
	return utils.Sum(digits)
}

/**
 * P17() - Number Letter Counts
 *  If all the numbers from 1 to 1000 (one thousand) inclusive were
 * 		written out in words, how many letters would be used?
 *	Date	2025.02.01
 */
func P17() int64 {
	// generate list using my num2words package
	sum := int64(0)
	for i := 1; i <= 1000; i++ {
		str := n2w.Num2Words(inew(int64(i)), true)
		no_space := strings.ReplaceAll(str, " ", "")
		no_hyphen := strings.ReplaceAll(no_space, "-", "")
		sum += int64(len(no_hyphen))
	}
	return sum
}

/**
 * P18() - Find the maximum total from top to bottom of the triangle
 *  Date	2025.02.09
 */
func P18() int64 {
	raw := "75\n95 64\n17 47 82\n18 35 87 10\n20 04 82 47 65\n19 01 23 75 03 34\n88 02 77 73 07 63 67\n99 65 04 28 06 16 70 92\n41 41 26 56 83 40 80 70 33\n41 48 72 33 47 32 37 16 94 29\n53 71 44 65 25 43 91 52 97 51 14\n70 11 33 28 77 73 17 78 39 68 17 57\n91 71 52 38 17 14 91 43 58 50 27 29 48\n63 66 04 68 89 53 67 30 73 16 69 87 40 31\n04 62 98 27 23 09 70 98 73 93 38 53 60 04 23"

	// first, create the triangle (just a slice of slices)
	rows := strings.Split(raw, "\n")

	// create the triangle "grid"
	grid := make([][]int64, len(rows))
	for i := range grid {
		grid[i] = make([]int64, i+1)
	}

	// extract values from strings
	for i, str := range rows {
		rowvals := strings.Split(str, " ")
		for j, v := range rowvals {
			val, _ := strconv.ParseInt(v, 10, 64)
			grid[i][j] = val
		}
	}

	// now, compute the paths. This is a connected tree, not a binary tree!
	for i := len(grid) - 2; i >= 0; i-- {
		row := grid[i] // which row is being worked on
		for j := 0; j < len(row); j++ {
			left := float64(grid[i+1][j])
			right := float64(grid[i+1][j+1])
			row[j] += int64(math.Max(left, right))
		}
	}

	return grid[0][0] // return the root!
}

/**
 * P19() - How many Sundays fell on the first of the month
 *		during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
 *	Date	2025.02.09
 */
func P19() int64 {
	sum := int64(0)
	for y := 0; y < 100; y++ {
		for m := 1; m <= 12; m++ {
			t := time.Date(1901+y, time.Month(m), 1, 0, 0, 0, 0, time.UTC)

			if t.Weekday() == 0 { // 0 is Sunday
				sum++
			}
		}
	}
	return sum
}

/**
 * P20() Factorial Digit Sum
 * 	Find the sum of the digits in the number 100!.
 * Date		2025.02.09
 */
func P20() int64 {
	num := fact(inew(100))
	dig := utils.Digits(num.String())
	return utils.Sum(dig)
}
