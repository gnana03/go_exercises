/* Below should create a trade by comparing the values 
   if the values are invalid should show an error while creating */

package main

import (
    "fmt"
    )

type Trade struct{
    name string
    volume int
    price float64
    buy bool
}

func CreateNewTrade(name string,volume int,price float64,buy bool)(*Trade,error) {
    if name==""{
       return nil,fmt.Errorf("Name cannot be empty")
     }
    if volume <=0 {
       return nil,fmt.Errorf("Volume cannot be zero or negative")
     }
    if price <=0.0 {
      return nil,fmt.Errorf("Price cannot be zero or negative")
    }
    trade := &Trade{name:name,volume:volume,price:price,buy:buy}
    return trade,nil
     }

func (t Trade) get_share_value() float64 {
    val := float64(t.volume)*t.price
    if t.buy {
       val=-val
   }
   return val
}

func main() {
   t,err := CreateNewTrade("AMAZON",123,1999.9,true)
   if err!=nil {
      fmt.Printf("Error is : %s \n",err)
   } else {
     fmt.Printf("Trade created is %+v ",t)
     fmt.Println("Share price is ",t.get_share_value())
   }
   t1,err := CreateNewTrade("",123,199.8,false)
   if err!=nil {
      fmt.Printf("Error is : %s \n",err)
   } else {
     fmt.Printf("Trade created is %+v ",t1)
    fmt.Println("Share price is ",t1.get_share_value())
  }
   t2,err := CreateNewTrade("MSFT",-123,199.8,false)
   if err!=nil {
      fmt.Printf("Error is : %s \n",err)
   } else {
     fmt.Printf("Trade created is %+v ",t2)
    fmt.Println("Share price is ",t2.get_share_value())
 }
   t3,err := CreateNewTrade("MSFT",100,-199.8,false)
   if err!=nil {
      fmt.Printf("Error is : %s \n",err)
   } else {
     fmt.Printf("Trade created is %+v ",t3)
    fmt.Println("Share price is ",t3.get_share_value())
 }
  
}
