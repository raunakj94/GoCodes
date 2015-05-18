package main
import "fmt"
import "time"
func simulateEvent(D string, timeInSecs int64){ 
fmt.Println(D "starts now and finish in " timeInSecs "seconds")
time.Sleep(time.Duration(timeInSecs * 1e9))
fmt.Println("finished",D)
}
func main(){
simulateEvent("race",10)
simulateEvent("Long Jump",6)
simulateEvent("High Jump",3)
}

