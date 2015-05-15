package main
import "fmt"
func main(){
a:=map[string]map[string]int{
"rahul":=map[string]int{
"age":66,
"salary":56890,
},
"ravi":=map[string]map[string]int{
"age":34,
"salary":567656,
},
}
if temp,  hero:= a["ravi"]; hero {
fmt.Println(temp["age"], temp[salary])
}
}



