// Go has garbage collector
// When we allocate some resource and dont use it GO's garbage collector cleans it
// defer will be called at last to ensure the cleaning is called 
//defer will be called even if there is any error in the code

package main
import (
    "fmt"
  )

func cleanup(name string){
  fmt.Println("Cleaning:  ",name)
}

func worker(){
   fmt.Println("Worker function")
   defer cleanup("A") // this is how we defer
   defer cleanup("B") // this would be called first
   fmt.Println("Worker")
}

func main() {
   worker()
}
