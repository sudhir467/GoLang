/* This program is to print the following pattern
*****
*****
*****
*****
*****
 */
package main

import (
	"fmt"
)

func main() {
	var rows, cols, num int
	fmt.Printf("Enter the number of rows required:")
	fmt.Scanf("%d", &num)
	for rows = 1; rows <= num; rows++ {
		for cols = 1; cols <= num; cols++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
