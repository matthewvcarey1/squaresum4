package squares

import (
	"fmt"
	"runtime"
)

type Squares struct {
	sqs        []int64
	size       int
	actualSize int
}

type Squareser interface {
	FindSumsOfSquares(ch chan string)
}

type binchopin struct {
	m  int64
	mi int
	k int64
	ki int 
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

// binChop calling sort.Search is slower than hand built binChop
// by 4.8021043 seconds against 2.6710666 seconds for a size of 10000
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
func (sq Squares) runsearch(inch chan binchopin, ch chan string, stopchan chan bool){
	var indata binchopin
	dostop := false
	for{
		dostop =<- stopchan
		if dostop {
			return
		} 
		indata = <-inch
		bottom := indata.mi + indata.ki                    // The index of the larger value
		top := (((indata.mi + indata.ki + 1) * 142) / 100) // The result index if both values were the same
		value := indata.m + indata.k
	
		res, ok, index := sq.binChop(value, bottom, top)
		
		if ok {
		// roots are the index values plus 1 as we started the array at 1
			kroot := indata.mi + indata.ki + 1
			mroot :=indata. mi + 1
			root := index + 1
			ch <- fmt.Sprintf("%d (%d * %d) + %d (%d * %d) = %d (%d * %d)",
				indata.m, mroot, mroot, indata.k, kroot, kroot, res, root, root)
		}
	}
}

func (sq Squares) FindSumsOfSquares(ch chan string) {
	defer close(ch)
	inch := make (chan binchopin)
	stopchan := make (chan bool)
	numCPUs := runtime.NumCPU()
	for i :=0; i<numCPUs;i++{
		go sq.runsearch(inch,ch, stopchan)
	}
	for mi, m := range sq.sqs[:sq.size] {
		// We start the inner loop from current outer loop index so that
		// we don't find duplicate commutative sums
		// eg: 9 + 16 = 16 + 9 = 25
		for ki, k := range sq.sqs[mi:sq.size] {
			indata := binchopin{ 
				m: m,
				mi: mi,
				k: k,
				ki: ki,
			}
			stopchan <- false
			inch <- indata
		}
		
	}
	for i :=0; i<numCPUs; i++{
		stopchan <- true
	}

}
