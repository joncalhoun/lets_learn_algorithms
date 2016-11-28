package main

import "fmt"

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
	bubbleSort(users)
	fmt.Println("Sorted:  ", users)
}

func bubbleSort(users []User) {
	// Implement this and any other functions you may need
	// If you are writing Go code, you can access the first
	// and last name like so:
	fmt.Println(users[0].FirstName)
	fmt.Println(users[0].LastName)
}
