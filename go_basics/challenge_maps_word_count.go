// this should lower the words and count the words in a string
// use strings.Fields to split the words
// use strings.ToLower to change to lower case

package main

import (
   "fmt"
   "strings"
   )

func main() {
    text := `
        Needles and pins
        Needles and Pins
        Sew me sail
        To catch me the wind`
  words  := strings.Fields(text) 
  words_count := map[string]int {}
  for _,word:= range words {
       words_count[strings.ToLower(word)]++
    }
  fmt.Println(words_count)
}
    
