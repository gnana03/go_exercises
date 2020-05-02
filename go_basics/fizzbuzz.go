/* if number is divisible by 3 print fizz
   if number is divisible by 5 print buzz
   if number is divisible by both 3 and 5 print fizzbuzz instead of the number */

package main

import (
   "fmt"
)

func main() {
  n:=20
  for i:=1;i<=n;i++ {
     switch {
        case i%3==0 && i%5==0:
         fmt.Println("fizzbuzz")
        case i%3==0:
         fmt.Println("fizz")
        case i%5==0:
         fmt.Println("buzz")
        default:
         fmt.Println(i)
     }
}
}
