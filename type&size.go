/* 
  This is the program for finding the type and size of the number.
  Created By: Mani Agarwal
  Email : agarwalmani22@gmail.com
*/

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a, b := 89,95
	fmt.Println("Value of a is ",a , "& b is ", b)
	fmt.Printf("Type of a is %T, Size of a is %d ", a, unsafe.Sizeof(a)) 
	fmt.Printf("\nType of b is %T, Size of b is %d ", b, unsafe.Sizeof(b)) 
}
