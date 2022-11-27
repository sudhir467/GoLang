/*Program to Find whether number is palindrome or not*/

package main

// Import the fmt package
import (
	"fmt"
)

func main() {
	var num, remainder, check int
	reverse := 0

	fmt.Println("Enter the number:")
	fmt.Scanf("%d", &num)
	check = num
	// while the number is not equal to zero compute the reverse of that number
	for num != 0 {
		remainder = num % 10
		reverse = reverse*10 + remainder
		num = num / 10
	}
	// check whether the number is equal to the reverse of the number and print the result
	if check == reverse {
		fmt.Println("The number is a Palindrome Number")
	} else {
		fmt.Println("The number is not a Palindrome Number")
	}

}
