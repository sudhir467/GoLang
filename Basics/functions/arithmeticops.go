/* This program is to calculate arithematic operations*/

package main

import "fmt"

func addition(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func subtraction(num1 int, num2 int) int {
	sub := num1 - num2
	return sub
}

func multiplication(num1 int, num2 int) int {
	mul := num1 * num2
	return mul
}

func main() {

	add := addition(21, 13)
	sub := subtraction(34, 14)
	mul := multiplication(5, 6)
	fmt.Println("sum:", add)
	fmt.Println("sub:", sub)
	fmt.Println("mul:", mul)

}
