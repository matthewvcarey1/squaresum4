package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/matthewvcarey1/squaresum4/internal/squares"
)

// Find squares that are sums of two other squares by brute force without using
// square root math or floating points.

// We create and ordered array of the square numbers then we loop within a loop
// through the lower part of the array and see if the sum of two of the squares
// are in the array.

// The returned index of the found square number tells us the square root of the number.

// A number passed the command line is taken as the query range of square roots

// The efficiency of this program should be Big O n^2 log n

func main() {
	t1 := time.Now()
	defer func() {
		taken := time.Now().Sub(t1)
		fmt.Println("Time taken", taken)
	}()
	size := 100
	if len(os.Args) > 1 {
		sizeStr := os.Args[1]
		ns, err := strconv.Atoi(sizeStr)
		if err == nil && ns > 0 {
			size = ns
		}
	}
	sq := squares.New(size)
	ch := make(chan string)

	go sq.FindSumsOfSquares(ch)
	for msg := range ch {
		fmt.Println(msg)
	}
}
