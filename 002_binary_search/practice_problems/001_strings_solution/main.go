package main

import (
	"fmt"
	"strings"
)

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
	var lo int = 0
	var hi int = len(words) - 1

	for lo <= hi {
		var mid int = lo + (hi-lo)/2
		var midValue string = words[mid]

		if compare(midValue, word) == 0 {
			return mid
		} else if compare(midValue, word) > 0 {
			// We want to use the left half of our list
			hi = mid - 1
		} else {
			// We want to use the right half of our list
			lo = mid + 1
		}
	}

	// If we get here we tried to look at an invalid sub-list
	// which means the number isn't in our list.
	return -1
}

func compare(a, b string) int {
	var aLow string = strings.ToLower(a)
	var bLow string = strings.ToLower(b)
	if aLow == bLow {
		return 0
	} else if aLow < bLow {
		return -1
	} else {
		return 1
	}
}
