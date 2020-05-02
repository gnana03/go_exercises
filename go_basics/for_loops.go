// for loops

package main

import (
    "fmt"
   )

func main() {
   for i:=0;i<5;i++ {
     fmt.Println(i)
   }

 // using break
fmt.Println("---break---")
for i:=0;i<5;i++ {
   if i>1 {
     break
   }
     fmt.Println(i)
   }

// using continue
fmt.Println("--- continue ---")
for i:=0;i<5;i++ {
  if i<1 {
   continue
   }
   fmt.Println(i)
  }

//like a while loop
fmt.Println("--- while loop--- ")
a:=0
for a<5 {
 fmt.Println(a)
 a++
}

// like a do while loop
fmt.Println("--- do while loop ---")
b:=0
for{
  if b>2 {
    break
}
 fmt.Println(b)
 b++
}

}
     

//
