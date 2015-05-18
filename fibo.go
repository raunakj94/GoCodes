package main
import "fmt"
func main(){
ch:=make(chan int)
go fib(ch)
for i:=0 ; i<20; i++ {
     
			fmt.Println(<-ch)
}
}
func fib(ch chan int){
i,j := 0,1
for{
ch<-j
i,j=j,i+j
}
}

