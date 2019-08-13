package main

import "fmt"

func addtwox(x uint64) {
	for x < 10000000000 {
		fmt.Println(x)
		x += 2 * x
	}
}

func main() {
	// will execute 3, 4, 1, 2
	defer addtwox(2)
	defer addtwox(1)
	addtwox(3)
	addtwox(4)
}
