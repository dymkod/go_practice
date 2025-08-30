package main

import "fmt"

func main() {
	var n, prod int = 0, 1
	var a []int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		a = append(a, x)
	}

	for i := 0; i < len(a); i++ {
		// of course we can do it on reading stage, its just a practice here
		prod *= a[i]
	}

	fmt.Printf("%d", &prod)
}
