/* heaps.go
   Implementation of heaps algorithm for permutations.
   
   Copyright (c) 2019, Michael Pascale
*/

package main

import "os"
import "fmt"
import "strconv"

package main

import "os"
import "fmt"
import "strconv"

func generate(n int, A []int) {

        if n == 1 {
                fmt.Println(A)
        } else {
                for i := 0; i < n-1; i++ {
                        generate(n-1, A)

                        var swap int
                        if n%2 == 0 {
                                swap = A[i]
                                A[i] = A[n-1]
                                A[n-1] = swap
                        } else {
                                swap = A[0]
                                A[0] = A[n-1]
                                A[n-1] = swap
                        }
                }
                
                generate(n-1, A)
        }
}

func main() {

        n, err := strconv.Atoi(os.Args[1]) // Get n from args...

        if err != nil {    // Atoi returns the value and an error code
                panic(err) // Check for an error code...
        }

        var A = make([]int, n)   // Allocate an array

        for i := 0; i < n; i++ { // Assign each value of the array a number
                A[i] = i + 1
        }
        
        generate(n, A)           // Run heap's algorithm
}
