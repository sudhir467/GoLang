/*
This program is to print the following pattern
11111
22222
33333
44444
55555
*/
package main

import (
	"fmt"
)

func main() {
	var rows, cols, number int

	fmt.Printf("Enter the number of rows required:")
	fmt.Scanf("%d", &number)
	for rows = 1; rows <= number; rows++ {
		for cols = 1; cols <= number; cols++ {
			fmt.Printf("%d", rows)
		}
		fmt.Printf("\n")
	}

}
