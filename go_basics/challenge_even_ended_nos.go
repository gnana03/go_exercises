/* Even ended numbers are the numbers having same first and last digit number
  example 11,121,232 etc.,
  Check wheher a number is an even ended number or not
  Check how many numbers even numbers result from multiplying  two four digit number 
   example 1011 * 1001 = 1012011 is even ended */

/* It is easier to find the even numbered numbers by converting
   This can be done using Sprintf function from fmt */

package main

import (
   "fmt"
  )

func main() {
  i:=42
  s:=fmt.Sprintf("%d",i)
  
  fmt.Printf("s=%q,type of %T\n",s,s)
  if s[0]==s[len(s)-1] { 
    fmt.Println("It is an even ended number")
   }else{
    fmt.Println("It is not an even ended number")
  }  
 count:=0
 for a:=1000;a<=9999;a++ {
   for b:=a;b<=9999;b++ {
      c := a*b
      str:= fmt.Sprintf("%d",c)
      if str[0] == str[len(str)-1] {
        count++;
     }
    }
  }
 fmt.Println("Number of even ended numbers are ",count)
 
}
