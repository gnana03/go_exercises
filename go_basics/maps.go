// Maps a data structure with keys and values

package main
import (
    "fmt"
)

func main() {
 stocks := map[string]float64 {
            "AMAZON":1666.6,
             "GOOG":1118.2,
             "MSFT":112, // Must have trailing comma till last line
              }
 fmt.Println("length of stocks",len(stocks))
 fmt.Println("Amazon stock price",stocks["AMAZON"])
 fmt.Println("Should give a zero value if not found",stocks["TSLA"])
 //set value
 stocks["TSLA"]=1122.3
 fmt.Println("After setting the value",stocks["TSLA"])
 delete(stocks,"MSFT")
 fmt.Println("After deleting MSFT fetching the value",stocks["MSFT"])
 // getting keys
 for key:= range stocks {
    fmt.Println(key)
}
// accessing both key value pair
for key,value:= range stocks {
   fmt.Printf("%s --> %.2f \n ",key,value)
}
}
