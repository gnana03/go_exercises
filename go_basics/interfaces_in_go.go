/* In Go language, the interface is a custom type that is used to specify a set of one or more method signatures and the interface is abstract, so you are not allowed to create an instance of the interface. But you are allowed to create a variable of an interface type and this variable can be assigned with a concrete type value that has the methods the interface requires. Or in other words, the interface is a collection of methods as well as it is a custom type.*/

/* In the Go language, it is necessary to implement all the methods declared in the interface for implementing an interface. The go language interfaces are implemented implicitly. And it does not contain any specific keyword to implement an interface just like other languages. As shown in the below example */

package main

import (
    "fmt"
    "math"
    )

type Shapes interface {
    Area() float64
}

type Square struct{
     Length float64
}

type Circle struct{
    Radius float64
}

type Rectangle struct{
    x float64
    y float64
}

func (s *Square) Area() float64 {
     return s.Length*s.Length
}

func (c *Circle) Area() float64 {
    return math.Pi*c.Radius*c.Radius
}

func (r *Rectangle) Area() float64 {
    return r.x*r.y
}

func SumOfAreas(shapes []Shapes) float64 {
   total :=0.0
   for _,shape := range shapes{
      total +=shape.Area()
}
return total
}

func main() {
 s := &Square{10.0}
 c := &Circle{5.0}
 r := &Rectangle{4.0,5.0}
 sa := s.Area()
 ca := c.Area()
 ra := r.Area()
 fmt.Println("Area of Square is ",sa)
 fmt.Println("Area of Cictle is ",ca)
 fmt.Println("Area of Rectangle is ",ra)
 shapes := []Shapes{s,c,r}
 sum_areas:= SumOfAreas(shapes)
 fmt.Println("Sum of all areas is ",sum_areas)
}



     
