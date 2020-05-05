// Like in any other languages such as Python,Java or C errors should be handled in order not to stop the execution abruptly
// For this to achieve we need to install github.com/pkg/errors using go get github.com/pkg/errors
// Below example shows what happens when the errors are not handled properly

package main
import (
   "fmt"
   "os"
   "log" // to stacktrace logs
   "github.com/pkg/errors"
)

type Config struct {
    // TBD
   }

func ReadConfig(filepath string)(*Config,error){
    file,err := os.Open(filepath)
    if err!=nil {
       return nil,errors.Wrap(err,"Error : Can't Open the file")
    }
    defer file.Close()
    cfg := &Config{}
    return cfg,nil
}
func setUpLogging() {
    out,err := os.OpenFile("app.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err !=nil {
        return
    }
    log.SetOutput(out)
}

func main() {
   setUpLogging()
   filepath := "/path/to/config.yaml"
   config,err := ReadConfig(filepath)
   if err!=nil {
      fmt.Fprintf(os.Stderr, "error : %s \n",err)
      log.Printf("error : %+v",err)
      os.Exit(1)
    }
   fmt.Println(config)
    }
      
