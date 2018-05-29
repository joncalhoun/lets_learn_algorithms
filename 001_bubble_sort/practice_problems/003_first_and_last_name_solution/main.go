package main

import (
	"fmt"
	"strings"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	var users []User = []User{
		User{
			FirstName: "Jon",
			LastName:  "Calhoun",
		},
		User{
			FirstName: "Bob",
			LastName:  "Smith",
		},
		User{
			FirstName: "John",
			LastName:  "Smith",
		},
		User{
			FirstName: "Larry",
			LastName:  "Green",
		},
		User{
			FirstName: "Joseph",
			LastName:  "Page",
		},
		User{
			FirstName: "George",
			LastName:  "Costanza",
		},
		User{
			FirstName: "Jerry",
			LastName:  "Seinfeld",
		},
		User{
			FirstName: "Shane",
			LastName:  "Calhoun",
		},
		User{
			FirstName: "Alice",
			LastName:  "Calhoun",
		},
	}

	fmt.Println("Unsorted:", users)

	for i := 0; i < len(users); i++ {
		bubbleSort(users[:len(users)-i])
	}

	fmt.Println("Sorted:  ", users)

}

func bubbleSort(users []User) {
	fmt.Println(users[0].FirstName)
	fmt.Println(users[0].LastName)

	var lngth int = len(users)
	var indx1 int = 0
	var indx2 int = 1

	for indx2 < lngth {
		var val1 User = users[indx1]
		var val2 User = users[indx2]

		if asc(users[indx1], users[indx2]) {
			users[indx1] = val2
			users[indx2] = val1
		}

		indx1++
		indx2++
	}

}

func asc(first User, second User) bool {
	if strings.ToLower(first.LastName + first.FirstName) < strings.ToLower(second.LastName + second.FirstName) {
		return true
	}
	return false
}
