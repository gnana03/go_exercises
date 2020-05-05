/* This challenge should do below things
  Write a function that gets a URL and returns the value of content type response header
  The function should return an error if can't perform the GET request */

package main
import (
	"fmt"
	"net/http"
)

func get_http_content_type(url string) (string,error) {
   resp,err := http.Get(url)
   if err != nil {
     return " ",err
  } else {
   ctype := resp.Header.Get("Content-Type")
   if ctype == "" { // should return error if content type not found
      return "",fmt.Errorf("Cannot find the content type")
   }
    defer resp.Body.Close()
   return ctype,nil
 }
}
func main() {
  url := "https://golang.org"
  ctype,err := get_http_content_type(url)
  if err!= nil {
    fmt.Printf("Error is %s",err)
  } else {
  fmt.Println(ctype)
 }  
}


