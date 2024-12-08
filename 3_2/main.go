package main

import (
	"fmt"
)

func insertionSort(A []int, N int) {
	for i := 1; i < N; i++ {
		v := A[i]      
		j := i - 1   

		for j >= 0 && A[j] > v {
			A[j+1] = A[j]
			j--
		}

		A[j+1] = v

		printArray(A)
	}
}

func printArray(A []int) {
	for i, v := range A {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func main() {
	var N int

	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	printArray(A)

	insertionSort(A, N)
}
