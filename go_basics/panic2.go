// This shows how we recover from panic it is just like handling the exceptions in other languages

package main

import (
   "fmt"
   )

func get_indexed_val(list []int,index int) int {
  defer func() {
    if err := recover();err!=nil {
       fmt.Printf("Error is %s \n",err)
     }
   }()
   return list[index]
 }

func main() {
    list := []int {1,2,3}
    v:=get_indexed_val(list,10)
    fmt.Println(v)
}
