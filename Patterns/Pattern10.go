/* This program is to print the following pattern
    1
   12
  123
 1234
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
			if rows+cols <= num {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%d", rows+cols-num)
			}
		}
		fmt.Println()
	}
}
