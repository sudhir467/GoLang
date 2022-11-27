/* This program is to find if the number is armstrong number or not*/

package main

// Import fmt package
import (
	"fmt"
)

// Main Function to compute whether number is Armstrong number or not
func main() {
	var remainder, num, check int
	cube := 0
	sum := 0

	fmt.Println("Enter the number to find if its Armstrong number or not:")
	fmt.Scanf("%d", &num)
	check = num // Intialize the check var with number given

	// while the number is not equal to zero compute whether Armstrong or not i.e.
	//sum of cubes of individual digits is equal to the number given
	for num != 0 {
		remainder = num % 10
		cube = remainder * remainder * remainder
		sum = sum + cube
		num = num / 10
	}
	// check If the number given is equal to the sum of individual cubes and then print if Armstrong or not
	if check == sum {
		fmt.Println("The number is Armstrong Number")
	} else {
		fmt.Println(" The number is not Armstrong Number")
	}

}
