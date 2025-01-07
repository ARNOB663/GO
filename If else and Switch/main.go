package main

import "fmt"

func main() {
    age := 18
    if age >= 18 { // If statement
        fmt.Println("You can vote") // Output: You can vote
    } else {
        fmt.Println("You can't vote") // Output: You can't vote
    }
	// Output: You can vote

	// Switch
	switch age {
	case 18:
		fmt.Println("You are 18 years old")
	case 19:
		fmt.Println("You are 19 years old")
	default:
		fmt.Println("You are not 18 or 19 years old")
	}
}