// This code should give the maximal value from the slice

package main

import (
   "fmt"
   )

func main() {
    
   nums:=[]int{16,18,289,27,3,0,222}
   max:=nums[0]
   for i:=0;i<len(nums);i++ {
      if nums[i]>max {
          max=nums[i]
      }
   }
   fmt.Println("maximum value from the list is",max)

// this can be changed with below code
max_value:=nums[0]
 for _,value:= range nums[0:] {
   if value>max_value {
      max_value=value
   }
 }
   fmt.Println("maximum value from the list is",max_value) 
}
