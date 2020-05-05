// This shows example of functions

package main
import (
    "fmt"
    )

func add(a int,b int) int {   // this is how we take the arguments into functions
     c := a+b
     return c
}

func divide(a int,b int) (int,int) {
    return a/b,a%b // to return multiple values
}

func main() { // this is the main function to execute
   val := add(3,5)
   fmt.Println("result is ",val)
   div,mod := divide(8,3)
   fmt.Println("div and mod are : ",div,mod)
 }

