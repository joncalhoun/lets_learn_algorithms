package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numbers []int = rand.Perm(10)
	fmt.Println(numbers)
	bubbleSort(numbers)
	fmt.Println(numbers)
}

func bubbleUp(numbers []int, n int) bool {
	var swapped bool = false
	for i := 0; i < n-1; i++ {
		var first int = numbers[i]
		var second int = numbers[i+1]
		if first > second {
			numbers[i] = second
			numbers[i+1] = first
			swapped = true
		}
	}
	return swapped
}

func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		// Terminate early if bubbleUp doesn't swap anything, and reduce the max index we look at by len(numbers) - i
		if !bubbleUp(numbers, len(numbers)-i) {
			return
		}
	}
}
