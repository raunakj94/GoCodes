package main
import "fmt"
import "time"
import "math/rand"
func bad(msg string,c chan string){
for i:=0; ; i++{
 		c<- fmt.Sprintf("%s %d", msg, i)
        time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
}
}
func main(){
 c := make(chan string)
    go bad("boring!", c)
    for i := 0; i < 5; i++ {
        fmt.Printf("You say: %q\n", <-c)
    }
    fmt.Println("You're Dumb,i m leaving.")
}

