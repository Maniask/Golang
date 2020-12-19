/* 
  This is the program for understanding the type conversion in go.
  Created By: Mani Agarwal
  Email : agarwalmani22@gmail.com
*/

// Go is very strict about the type. There is no automatic type conversion.

package main

import (
	"fmt"
)

func main() {
	a, b := 55,67.8
	sum := a+ int(b)
	fmt.Println("Sum is ", sum) 
}
