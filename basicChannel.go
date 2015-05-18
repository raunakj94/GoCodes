package main
import (
"fmt"
"time"
)
func main(){
say("hello",3)
say("hiii",5)
 say("bye",2)
time.Sleep(8 * time.Second)
}
func say(msg string,secs int){
time.Sleep(time.Duration(secs) * time.Second)
fmt.Println(msg)
}
