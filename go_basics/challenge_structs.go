/* Define a Square struct with two fields : center of type point and length of type int
   Add two methods 
   1. Move(dx int,dy int)
   2. Area () int
  also write 
   3. NewSquare(x int,y int,length int)(Square,error)
*/

package main
import (
    "fmt"
     )

type Point struct {
    x int
    y int
   }

type Square struct {
    Center Point
    Length int
}

func (p *Point) Move (dx int,dy int) {
    p.x += dx
    p.y += dy
}

func (s *Square) Move (dx int,dy int) {
    s.Center.Move(dx,dy)
  }


func NewSquare (x int,y int,length int) (*Square,error) {
   if length <=0 {
     return nil,fmt.Errorf("Length cannot be less than or equal to zero")
   } else {
     s := &Square{
           Center: Point{x,y},
           Length: length,
           }
     return s,nil
     }
}
  
func (s1 *Square) Area() (int,error) {
      if s1.Length <= 0 {
           return 0,fmt.Errorf("Area cannot be calculated")
      } else {
            return s1.Length*s1.Length,nil
   }
 }


func main() {
    s,err := NewSquare(1,1,10)
    s.Move(2,3)
    area,err:=s.Area()     
    if err!=nil {
      fmt.Printf("Error is %s",err)
    } else {
      fmt.Printf("Area of square with length = %d is %d \n", s.Length,area)
      fmt.Printf("After moving the x and y new x and y are %+v",s)
}
}
 
    
