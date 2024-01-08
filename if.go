package main

import (
	"fmt"
)

func main() {
	// Replace this with the number you want to check
	number := 7

	if number%2 == 0 {
		fmt.Printf("%d is even.\n", number)
	} else {
		fmt.Printf("%d is odd.\n", number)
	}
}
