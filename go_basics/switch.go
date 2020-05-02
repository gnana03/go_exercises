// To verify switch statements

package main

import (
   "fmt"
 )

func main() {
  
  x:=11
  
  // To verify switch statements in integers
  
  switch x {
     case 1:
          fmt.Println("One")
     case 5:
         fmt.Println("Five")
     case 10:
         fmt.Println("Ten")
     default:
        fmt.Println("Other number")
   }

 // To verify strings
   
   str:="Go"
   switch str {
      case "c":
         fmt.Println("C language")
      case "c++":
        fmt.Println("c++ laguage")
      case "Go":
        fmt.Println("Go language")
      case "python":
        fmt.Println("Python language")
      default:
        fmt.Println("Other language")
    }

  // to verify conditions
    
  y:=4
  switch{
    case y>5:
       fmt.Println(">5")
    case y>10:
       fmt.Println(">10")
    case y==20:
       fmt.Println("=20")
    case y%4==0 && y%5==0:
       fmt.Println("divisible by 4 and 5")
    default:
       fmt.Println("no condition satisfied")
   }
}
