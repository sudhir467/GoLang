/* This program is to count the number of digits of Number*/

package main

//Import fmt Package
import (
	"fmt"
)

// Main Function to calculate the number of digits
func main() {
	var num int
	count := 0

	fmt.Println("Enter the number:")
	fmt.Scanf("%d", &num)
	//check whether the number is zero
	if num == 0 {
		count = 1

	} else {
		//for number greater that zero we increment the variable for every digit
		for num != 0 {
			num = num / 10
			count++
		}

	}
	//Print the value of count
	fmt.Println("The number of digits of number are:", count)

}
