// To know how to pass by reference

package main
import (
     "fmt"
   )

func double_at_index(values []int,index int){
     values[index] *= 2
}

func doubleVal(num int) {
     num *=2
}

func doublePtr(num *int) {
    *num *=2
}

func main() {

    list_values := []int {1,2,3,4}
    double_at_index(list_values,2)
    fmt.Println(list_values) // now should print new values because list was modified by function
    val := 10
    doubleVal(val)
    fmt.Println(val) // as the passing by value the value is not changed
    doublePtr(&val) // to pass by reference
    fmt.Println(val) // new value should be printed

}

 
