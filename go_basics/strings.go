// String

package main
import (
   "fmt"
   )

func main() {
   book:= "This is a Go Language Book"
   fmt.Printf("book[0]=%v is type %T\n",book[0],book[0])

   // string slicing
   fmt.Println(book[4:])

   fmt.Println(book[0:4])
   fmt.Println(book[:5])
   
   //concatination
   fmt.Println("yes"+book)
   
   // unicode characters
   const nihongo = "日本語"
   fmt.Println(nihongo)

}
