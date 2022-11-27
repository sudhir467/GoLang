/* This program is to check if the number is postive or negative*/

package main

// Import fmt Package
import (
	"fmt"
)

// function to calculate positive or negative of num
func main() {
	var num int

	fmt.Println("Enter the number:")
	fmt.Scanf("%d", &num)
	//check if the number is greater than, less than or equal to zero
	if num > 0 {
		fmt.Println("Entered number is positive")
	} else if num < 0 {
		fmt.Println("Entered number is negative")
	} else {
		fmt.Println("Entered number is zero")
	}

}
