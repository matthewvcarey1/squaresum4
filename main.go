package main

import (
	"fmt"
	"os"
	"strconv"
)

// Looks for the value in the ordered set by binary chopping the ordered set
// This is a Big O log n efficency method. (Where log is log2).

func binchop(value int64, set []int64, index int) (int64, bool, int) {
	for {
		size := len(set)
		// Not found case
		if size < 1 {
			return -1, false, -1
		}
		half := size / 2
		test := set[half]
		// Found case
		if test == value {
			return test, true, index + half
		} else {
			if test < value {
				// Look in top half
				set = set[half+1:]
				index = index + half + 1
			} else {
				// Look in bottom half
				set = set[:half]
			}
		}
	}
}

// Find squares that are sums of two other squares by brute force without using
// square root math or floating points.

// We create and ordered array of the square numbers then we loop within a loop
// through the lower part of the array and see if the sum of two of the squares
// are in the array.

// The returned index of the found square number tells us the square root of the number.

// A number passed the command line is taken as the query range of square roots

// The efficiency of this program is Big O n^2 log n

func main() {
	size := 100
	if len(os.Args) > 1 {
		sizeStr := os.Args[1]
		ns, err := strconv.Atoi(sizeStr)
		if err == nil && ns > 0 {
			size = ns
		}
	}

	ch := make(chan string)
	go findSumOfSquares(size, ch)
	for msg := range ch {
		fmt.Println(msg)
	}
}

func findSumOfSquares(size int, ch chan string) {
	// Build an array of the square number big enough to hold
	// possible sums of the queried range "size"
	// The square root of sum of two squares of any one number
	// is number * 1.414213562... (an irrational number).
	// We will approximate this in integer maths to 1.42 giving
	// a little room for leaway but not calculating a vast number
	// of unused squares.
	arrsz := (size * 142) / 100
	sqs := make([]int64, arrsz)
	for n := 1; n <= arrsz; n++ {
		sqs[n-1] = int64(n) * int64(n)
	}

	for mi, m := range sqs[:size] {
		// We start the inner loop from current outer loop index so that
		// we don't find duplicate commutative sums
		// eg: 9 + 16 = 16 + 9 = 25
		for ki, k := range sqs[mi:size] {
			// try to look up the sum of squares in the array of squares
			// we are going to restrict the array to broadly possible indices
			bottom := mi + ki                    // The index of the larger value
			top := (((mi + ki + 1) * 142) / 100) // The result index if both values were the same
			res, ok, index := binchop(m+k, sqs[bottom:top], bottom)
			//res, ok, index := binchop(m+k, sqs, 0)
			// On success report
			if ok {
				// roots are the index values plus 1 as we started the array at 1
				kroot := mi + ki + 1
				mroot := mi + 1
				root := index + 1
				ch <- fmt.Sprintf("%d (%d * %d) + %d (%d * %d) = %d (%d * %d)",
					m, mroot, mroot, k, kroot, kroot, res, root, root)
			}
		}
	}
	close(ch)
}
