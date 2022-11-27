/* This program is to reverse a given number*/

package main

// Import the package fmt
import (
	"fmt"
)

// Main function of the program
func main() {
	var num, remainder int
	reverse := 0
	fmt.Println("Enter the number to be reversed:")
	fmt.Scanf("%d", &num)

	//while number is not equal to zero compute the reverse of the number
	for num != 0 {
		remainder = num % 10
		reverse = reverse*10 + remainder
		num = num / 10
	}
	// Print the reverse of the number
	fmt.Println("The reverse of the given number is:", reverse)
}
