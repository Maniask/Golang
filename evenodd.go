/* 
  This is the program for finding if the given number is even or odd .
  Created By: Mani Agarwal
  Email : agarwalmani22@gmail.com
*/

package main

import "fmt"

func main(){
    var num int
    fmt.Println("Enter the number")
    fmt.Scan(&num)
    if num%2==0{
        fmt.Println("Number is even")
    }else{
        fmt.Println("Number is odd")
    }
}
