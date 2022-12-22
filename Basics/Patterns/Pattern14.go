/* This program is to print the following pattern
* * * * *
 * * * *
  * * *
   * *
    *
*/

package main

import (
	"fmt"
)

func main() {
	var rows, cols, num, spaces int
	fmt.Println("Enter the number of rows required:")
	fmt.Scanf("%d", &num)
	for rows = 1; rows <= num; rows++ {
		for spaces = 1; spaces <= rows-1; spaces++ {
			fmt.Printf(" ")
		}
		for cols = 1; cols <= num-rows+1; cols++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
