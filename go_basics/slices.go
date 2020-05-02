// slices of the string

package main
import (
   "fmt"
)

func main() {
  loons:= []string{"bugs","issues","code","source"}
  fmt.Printf("loons are %v , type is %T\n",loons,loons)
 
  //length
  fmt.Println("length is ",len(loons))

  // 0 indexing
  fmt.Println("0 index of loons is ",loons[0])
 
 // slicing
  fmt.Println("All values except 0th index is",loons[1:])
}
 
