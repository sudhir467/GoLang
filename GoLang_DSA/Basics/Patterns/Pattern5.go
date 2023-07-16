/*This program is to print the following pattern
1
22
333
4444
55555
*/

package main

import (
	"fmt"
)

func main() {
	var rows, cols, num int
	fmt.Println("Enter the number of rows required")
	fmt.Scanf("%d", &num)
	for rows = 1; rows <= num; rows++ {
		for cols = 1; cols <= rows; cols++ {
			fmt.Printf("%d", rows)
		}
		fmt.Println()
	}
}
