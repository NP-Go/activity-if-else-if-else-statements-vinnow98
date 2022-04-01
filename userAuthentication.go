package main

import "fmt"

func main() {
	var authenticatedUser = []string{"Robin", "John"}
	var name string
	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)
	for _, element := range authenticatedUser {

		if element == name {
			fmt.Println("Welcome!")
			break
		} else if name == ("Admin") {
			fmt.Println("Welcome Admin!")
			break
		} else {
			fmt.Println("Merry men!")
			break
		}

	}

}
