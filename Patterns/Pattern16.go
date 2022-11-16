/* This program is to print the following pattern
A B C D E
 A B C D
  A B C
   A B
    A
*/

package main

import (
	"fmt"
)

func main() {
	var rows, cols, spaces, num int
	fmt.Println("Enter the number of rows required:")
	fmt.Scanf("%d", &num)
	for rows = 1; rows <= num; rows++ {
		for spaces = 1; spaces <= rows-1; spaces++ {
			fmt.Printf(" ")
		}
		for cols = 1; cols <= num-rows+1; cols++ {
			fmt.Printf("%c ", 64+cols)
		}
		fmt.Println()
	}
}
