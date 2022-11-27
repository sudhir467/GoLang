/* This program is to print the greatest number amoung three numbers*/

package main

//Import fmt package
import (
	"fmt"
)

// function to calculate the greatest of three numbers
func main() {

	var a, b, c int

	fmt.Printf("Enter the values of a,b and c:")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	//check if a is greater than b and c
	if a > b && a > c {
		fmt.Println("a is greater than b and c ")

	} else if b > a && b > c { //check if b is greater than a and c

		fmt.Println("b is greater than a and c")

	} else if c > a && c > b { //check if c is greater than a and b
		fmt.Println("c is greater than b and a")
	}

}
