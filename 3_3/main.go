package main

import "fmt"

// 5 3 2 4 1
// 3 2 4 1 5 → 2 3 4 1 5 → 2 3 1 4 5
func main() {
    var n int
    fmt.Scan(&n)
    a := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&a[i])
    }

    compareCount := 0 

    for i := 0; i < n-1; i++ { 
        for j := 0; j < n-1-i; j++ { 
            compareCount++ 
            if a[j] > a[j+1] {
                temp := a[j]      
                a[j] = a[j+1]    
                a[j+1] = temp     
            }
        }
    }

    for i, v := range a {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(v)
    }
    fmt.Println()
    fmt.Println(compareCount)
}
