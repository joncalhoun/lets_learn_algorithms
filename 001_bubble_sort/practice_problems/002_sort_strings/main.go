package main

import "fmt"

func main() {
	var animals []string = []string{
		"dog",
		"cat",
		"alligator",
		"cheetah",
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
	// Implement this and any other functions you may need
}
