package main

import (
	"fmt"
	"math"
)

const MAX = 200000

func main() {
	var n int
	var R [MAX]int 

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&R[i])
	}

	maxv := math.MinInt64 
	minv := R[0]         

	for i := 1; i < n; i++ {
		if profit := R[i] - minv; profit > maxv {
			maxv = profit
		}

		if R[i] < minv {
			minv = R[i]
		}
	}

	fmt.Println(maxv)
}
