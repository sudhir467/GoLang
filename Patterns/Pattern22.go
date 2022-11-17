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
			if rows == 1 {
				fmt.Printf("%d ", cols)
			} else {
				fmt.Printf("%d ", cols+rows-1)
			}
		}
		fmt.Println()
	}
	for rows = 1; rows < num; rows++ {
		for spaces = 1; spaces <= num-rows-1; spaces++ {
			fmt.Printf(" ")
		}
		for cols = 1; cols <= num-spaces+1; cols++ {
			if rows == 5 {
				fmt.Printf("%d ", cols)
			} else {
				fmt.Printf("%d ", cols+num-rows-1)
			}
		}
		fmt.Println()
	}
}
