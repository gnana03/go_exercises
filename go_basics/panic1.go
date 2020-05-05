// panic show the exact error in go

package main
import (
   "fmt"
   "os"
  )

func main() {
    file,err := os.Open("no-such-file")
    if err != nil {
       panic(err)
    }
    defer file.Close()
    fmt.Println("file is opened")
  }

