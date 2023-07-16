/* Calculate the area of rectangle*/

package main

import "fmt"

func rectangle(length int, width int) int {
	area := length * width
	return area
}

func main() {

	fmt.Println("Area of the rectangle:", rectangle(12, 13))
}
