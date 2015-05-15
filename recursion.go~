package main
import "fmt"
 func main() {
a:= []int{1,2,3,4,5}
fmt.Println("Sum :", addThemUp(a))
fmt.Println("Sub:",subThem(1,2,3,4,5,6,7))
fmt.Println("Mul:",mul(1,2,3,4,5,10))
fmt.Println("Multi;",multi(a))
}
func addThemUp(num []int) int{
sum := 0
for _, val := range num {
sum += val
}
return sum
}
func subThem(args ...int) int{
 
finalValue := 0
for _, value := range args {
finalValue -= value
}
return finalValue
 }
func mul(args ...int) int {
mul:=1
for _, val :=range args{
mul*=val
}
return mul
}

func multi(abc []int) int{
mul:=1
for _,val := range abc{
mul*=val
}
return mul
}


