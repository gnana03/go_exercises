/* Structs to combine several different types of data types
  example type struct1 struct {
             name string
             rollno int
             gpa float64
             pass bool
      }
 to access  one of the variable in struct use struct1.name or struct1.rollno etc
 to initialize struct use s1 := struct1{"ABC",123,7.8,true} or also as below
  s2 := struct1{
             name: "DEF"
             gpa: 4.0
             pass: false 
             rollno: 1234
     }
 Go will assign 0 value if not defined or if we use as below
  s3: = struct1{}
*/
package main
import (
   "fmt"
   )

type Trade struct {
            name string
            volume int
            share_price float64
            trade bool
         }
func main() {
     t1 := Trade{"AMAZON",100,1999.9,true}
     fmt.Printf("%+v\n",t1)
     t2 := Trade{
           name: "MSFT",
           trade: true,
           share_price: 1122.3,
           volume: 120,
      }
     fmt.Println(t2.name)
     fmt.Printf("%+v\n",t2)
     t3:=Trade{}
     fmt.Printf("%+v\n",t3)
   }

  
