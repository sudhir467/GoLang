/*This program is to print the following pattern
    *
   **
  ***
 ****
*****
 ****
  ***
   **
    *
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
		for spaces = 1; spaces <= num-rows; spaces++ {
			fmt.Printf(" ")
		}

		for cols = 1; cols <= num-spaces+1; cols++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	for rows = 1; rows < num; rows++ {
		for spaces = 1; spaces <= rows; spaces++ {
			fmt.Printf(" ")
		}
		for cols = 1; cols <= num-rows; cols++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
