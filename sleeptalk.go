package main
import (
"fmt"
"time"
)
func SaT(secs time.Duration,msg string){
time.Sleep(secs*time.Second)
fmt.Printf("%v",msg)
}
func main(){
go SaT(5,"hello")
go SaT(3,"hi whats up?")
go SaT(2,"yes")
time.Sleep(4 * time.Second)
}
