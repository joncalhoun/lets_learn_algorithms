package main

import (
	"fmt"
	"strings"
)

func main() {
	var animals []string = []string{
		"dog",
		"cat",
		"alligator",
		"Cheetah",
		"rat",
		"moose",
		"cow",
		"mink",
		"porcupine",
		"dung beetle",
		"camel",
		"steer",
		"bat",
		"hamster",
		"horse",
		"colt",
		"bald eagle",
		"frog",
		"rooster",
	}

	fmt.Println("Unsorted:", animals)
	bubbleSort(animals)
	fmt.Println("Sorted:  ", animals)
}

func bubbleSort(animals []string) {
	var N int = len(animals)
	var i int
	for i = 0; i < N; i++ {
		sweep(animals)
	}
}

func sweep(animals []string) {
	var N int = len(animals)
	var firstIndex int = 0
	var secondIndex int = 1

	for secondIndex < N {
		var firstVal string = animals[firstIndex]
		var secondVal string = animals[secondIndex]

		if greater(firstVal, secondVal) {
			animals[firstIndex] = secondVal
			animals[secondIndex] = firstVal
		}

		firstIndex++
		secondIndex++
	}
}

func greater(a, b string) bool {
	if strings.ToLower(a) > strings.ToLower(b) {
		return true
	} else {
		return false
	}
}
