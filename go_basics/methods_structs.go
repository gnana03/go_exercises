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
            buy bool
         }
func (t *Trade) get_share_value() float64 { // to pass the struct we should pass before the method name
    value := float64(t.volume)*t.share_price
    if t.buy{
      value=-value
    }
   return value
}
// below demonstrates the way to modify the struct
type Point struct {
   x int
   y int
}

func (p Point) Move(dx int,dy int) {
p.x += dx
p.y += dy
}

func (p *Point) MovePtr(dx int,dy int) {
  p.x += dx
  p.y += dy
}


func main() {
     t1 := Trade{"AMAZON",100,1999.9,true}
     fmt.Printf("%+v\n",t1)
     t2 := Trade{
           name: "MSFT",
           buy: true,
           share_price: 1122.3,
           volume: 120,
      }
     fmt.Println(t2.name)
     fmt.Printf("%+v\n",t2)
     t3:=Trade{}
     fmt.Printf("%+v\n",t3)
     fmt.Printf("The value of %s with volumes %d and share price %.2f is %.2f\n",t1.name,t1.volume,t1.share_price,t1.get_share_value()) // to call the method we use . operator
     p1:= Point{x:1,y:2}
     fmt.Println("Before Moving",p1)
     p1.Move(2,3)
     fmt.Println("After Moving",p1)
     // now to change the structs we have to use pointers as below
     p2 := &Point{x:2,y:3}
     fmt.Println("Before moving",p2)
     p2.MovePtr(3,4)
     fmt.Println("After Moving through pointer by 3 and 4 respectively",p2)
  }
