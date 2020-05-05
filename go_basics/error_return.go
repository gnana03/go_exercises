// To return error if error found

package main
import (
   "fmt"
   "math"
  )

func sqrt(n float64)(float64,error){
   if n<0 {
      return 0.0,fmt.Errorf("Square root of negative number %v cannot be calculated",n)
   } else {
    return math.Sqrt(n),nil // nil is same like null or none in any other language
   }
}
func main() {
    s1,err := sqrt(2.0)
    if err!= nil {
        fmt.Printf("Error is : %s \n",err)
    } else {
       fmt.Println("Square root of number is ",s1)
   }
   s2,err2 := sqrt(-2.0)
   if err2!= nil {
        fmt.Printf("Error is : %s \n",err2)
    } else {
       fmt.Println("Square root of number is ",s2)
   }
}
