/*
This program is to print the following pattern
1
2 3
4 5 6
7 8 9 10
*/
package main

import (
	"fmt"
)

func main() {
	var rows, cols, num int
	var count int = 1
	fmt.Println("Enter the number of rows required:")
	fmt.Scanf("%d", &num)
	for rows = 1; rows <= num; rows++ {
		for cols = 1; cols <= rows; cols++ {
			fmt.Printf("%d ", count)
			count++
		}
		fmt.Println()
	}

}
