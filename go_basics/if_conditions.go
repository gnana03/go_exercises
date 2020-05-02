// To check conditional statements

package main

import (
   "fmt"
    )

func main() {
   x := 10
   y := 20

   if x>5 {
     fmt.Printf("x=%v, is greater than 5\n",x)
   } else {
      fmt.Printf("x=%v, is not greater than 5\n",x)
    }
   
   if x%10==0 && y%10==0 {
       fmt.Printf("x=%v,y=%v are divisible by 10\n",x,y)
    } else {
      fmt.Println("x and y are not divisible by 10")
   }
 
  // check initialization and condition at the same time
   m:=10
   n:=5
   if frac:=m/n;frac>=2{
       fmt.Println("checked initialization and condition at the same time and found fracion>=0.5 which is %v",frac)

   } 
 }

