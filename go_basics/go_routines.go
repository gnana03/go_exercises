/* Concurrency is the composition of independantly executed processes can be achived through Go Routines
   example if we need to pass multiple URL values and get their content type then it would take lot of
   time to loop through and execute, this can be achieved through Go Routines */

/* To understand the difference uncomment below main functions and run time go run FILENAME and verify the
   time taken */
/* This can also be dealt with channels */

package main

import (
   "fmt"
   "net/http"
//   "sync"
  )
/*
func returnType(url string) {
    resp, err := http.Get(url)
    if err != nil {
         fmt.Printf("Error is %s",err)
         return
    }
    defer  resp.Body.Close()
    ctype := resp.Header.Get("content-type")
    fmt.Printf("%s --> %s \n",url,ctype)
  }
*/

func returnType(url string, out chan string) {
    resp, err := http.Get(url)
    if err != nil {
          out <- fmt.Sprintf("Error is %s",err)
         return
    }
    defer  resp.Body.Close()
    ctype := resp.Header.Get("content-type")
    out <- fmt.Sprintf("%s --> %s \n",url,ctype)
  }
// Below is the tradition approach which takes maximum time
/*
func main() {
   urls := []string {
               "https://golang.org",
                "https://api.github.com",
          }
   for _,url := range urls {
       returnType(url)
   }
*/
// Below uses go routines which simplifies time
/*
func main() {
   urls := []string {
               "https://golang.org",
                "https://api.github.com",
                "https://httpbin.org.xml", // ending element also should have ,
          }
   for _,url := range urls {
       go func(url string) {
             returnType(url)
       }(url)
   }
*/
// Above will not print anything because go runtime will not wait for go routines which are working
// This can be adressed by using sync packages
/*
func main() {
   urls := []string {
               "https://golang.org",
                "https://api.github.com",
          }
   var wg sync.WaitGroup
   for _,url := range urls {
       wg.Add(1)
       go func(url string) {
             returnType(url)
             wg.Done()
       }(url)
   }
   wg.Wait()
*/

func main() {
    urls := []string {
               "https://golang.org",
                "https://api.github.com",
          }
    ch := make(chan string)
    for _,url := range urls{
          go returnType(url, ch)
     }
     for range urls {
         out := <-ch
        fmt.Println(out)
     }


}


