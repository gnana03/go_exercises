/* Question : Read a pid file and verify the Process Id, convert to integer
   then do not kill the process Id but print saying to kill it 
   Use the error wrappers and strconv packages to do so */

package main
import (
   "fmt"
   "github.com/pkg/errors"
   "io/ioutil"
   "strconv"
   "os"
   "log"
   "strings"
   )

func killServer(pidFile string) error {
    data,err := ioutil.ReadFile(pidFile)
    if err!= nil {
      return errors.Wrap(err,"Cannot open the Pid File \n")
    }
    if err :=os.Remove(pidFile);err!=nil {
        log.Printf("Warnig: cant remove the pidFile \n")
     }
    strPID := strings.TrimSpace(string(data))
    pid,err := strconv.Atoi(strPID)
    if err!=nil {
       return errors.Wrap(err,"bad process Id") 
    }
    fmt.Printf("Killing server with pid : %d \n",pid)
    return nil
}

func main() {
   if err:=killServer("process.pid");err!=nil {
      fmt.Printf("There was an error: %s",err)
      os.Exit(1)
    }
    }
   
    
      

