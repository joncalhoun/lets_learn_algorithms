package main

import "fmt"

// Given a list of sorted words (strings with no spaces),
// search for a user provided word in the list without
// being case sensitive. Return -1 if the word isn't found
// and return the index of the word if it is found.
func main() {
	var words []string = []string{
		"ALLigator",
		"bat",
		"bEEtle",
		"camel",
		"cat",
		"cheetah",
		"COLT",
		"cow",
		"dog",
		"eagle",
		"froG",
		"hamster",
		"horse",
		"mink",
		"moose",
		"porcupine",
		"RaT",
		"rooster",
		"steer",
	}
	fmt.Println("Sorted Words:", words)
	var toFind string
	fmt.Println("What word should we search for? No spaces please!")
	fmt.Scanf("%s", &toFind)
	var index int
	index = binSearch(words, toFind)
	if index < 0 {
		fmt.Println("The word", toFind, "could not be found!")
	} else {
		fmt.Println("The word", toFind, "was found at index:", index, words[index])
	}
}

func binSearch(words []string, word string) int {
	return -1
}
