package  main
import "fmt"
func main(){
 a:=map[string]map[string]string{
  "Superman":map[string]string{
    "realname":"Clark Kent",
     "city":"metropolis",
},
"batman":map[string]string{
 "real":"Bruce",
"city" :"gotham",
},
}
if temp, hero := a["Superman"]; hero {
fmt.Println(temp["realname"], temp["city"])
}
}



 
