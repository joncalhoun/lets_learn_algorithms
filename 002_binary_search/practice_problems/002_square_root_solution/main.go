package main

import "fmt"

// Use a binary search to find the square root of a number
func main() {
	var number float64 = 128.0
	fmt.Println("Number:", number)
	var squareRoot float64 = squareRoot(number)
	fmt.Println("Square root:", squareRoot)
}

func squareRoot(n float64) float64 {
	lo := 0.0
	hi := n
	for i := 0; i < 2000; i++ {
		mid := lo + (hi-lo)/2.0
		if mid*mid > n {
			hi = mid
		} else {
			lo = mid
		}
	}
	return lo
}
