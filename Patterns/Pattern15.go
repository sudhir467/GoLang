/* This program is to print the following pattern
    1
   1 2
  1 2 3
 1 2 3 4
1 2 3 4 5
*/

package main

import (
	"fmt"
)

func main() {
	var rows, cols, num int
	fmt.Println("Enter the value of rows:")
	fmt.Scanf("%d", &num)
	for rows = 1; rows <= num; rows++ {
		for cols = 1; cols <= num; cols++ {
			if (rows + cols) <= num {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%d ", rows+cols-num)
			}
		}
		fmt.Println()
	}
}
