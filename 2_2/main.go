package main

import (
	"fmt"
	"sort"
)

func main(){ 
	var a []int = []int{25,36, 4, 55, 60, 60, 18 , 19, 78 , 49}
	sort.Slice(a, func(i, j int) bool{
		fmt.Printf("Comparing a[%d]=%d and a[%d]=%d\n", i, a[i], j, a[j])
		return a[i] > a[j]
	} ) 
	fmt.Println(a[0], a[1], a[2])
}


