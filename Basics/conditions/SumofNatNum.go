/* This program is to Print sum of natural numbers*/

package main

// Import fmt package
import (
	"fmt"
)

// function to calculate the sum
func main() {
	var num int
	sum := 0

	fmt.Println("Enter the number :")
	fmt.Scanf("%d", &num)
	// while the number is greater than zero we add individual digits
	//untill it it becomes zero

	for num > 0 {
		sum = sum + num
		num--
	}

	// Print the sum of all natural numbers
	fmt.Println("The sum of all natural numbers:", sum)
}
