package squares

import "fmt"

type Squares struct {
	sqs        []int64
	size       int
	actualSize int
}

type Squareser interface {
	FindSumsOfSquares(ch chan string)
}

func new(size int) Squares {
	arrsz := (size * 142) / 100
	sqs := make([]int64, arrsz)
	for n := 1; n <= arrsz; n++ {
		sqs[n-1] = int64(n) * int64(n)
	}
	return Squares{
		sqs:        sqs,
		size:       size,
		actualSize: arrsz,
	}
}

func New(size int) Squareser {
	return new(size)
}

func (sq Squares) binChop(value int64, start int, end int) (int64, bool, int) {
	set := sq.sqs[start:end]
	offset := start
	for {
		size := len(set)
		// Not found case
		if size < 1 {
			return -1, false, -1
		}
		half := size >> 1
		test := set[half]
		// Found case
		if test == value {
			return test, true, offset + half
		} else {
			if test < value {
				// Look in top half
				set = set[half+1:]
				offset = offset + half + 1
			} else {
				// Look in bottom half
				set = set[:half]
			}
		}
	}
}

func (sq Squares) FindSumsOfSquares(ch chan string) {
	defer close(ch)
	for mi, m := range sq.sqs[:sq.size] {
		// We start the inner loop from current outer loop index so that
		// we don't find duplicate commutative sums
		// eg: 9 + 16 = 16 + 9 = 25
		for ki, k := range sq.sqs[mi:sq.size] {
			// try to look up the sum of squares in the array of squares
			// we are going to restrict the array to broadly possible indices
			bottom := mi + ki                    // The index of the larger value
			top := (((mi + ki + 1) * 142) / 100) // The result index if both values were the same
			value := m + k
			res, ok, index := sq.binChop(value, bottom, top)
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
}
