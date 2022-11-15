/*
This program is to print the following pattern
12345
12345
12345
12345
12345
*/
package main

import (
	"fmt"
)

func main() {
	var rows, cols, num int
	fmt.Println("Enter the number of rows required:")
	fmt.Scanf("%d", &num)
	for rows = 1; rows <= num; rows++ {
		for cols = 1; cols <= num; cols++ {
			fmt.Printf("%d", cols)
		}
		fmt.Println()
	}

}
