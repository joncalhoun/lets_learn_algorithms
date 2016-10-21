package main

import "fmt"

func main() {
	var numbers []int = []int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}
	fmt.Println("Unsorted:", numbers)
	reverseBubbleSort(numbers)
	fmt.Println("Sorted:  ", numbers)
}

// This should sort numbers IN REVERSE ORDER!!!
func reverseBubbleSort(numbers []int) {
	// Implement this and any other functions you may need
}
