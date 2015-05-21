package main
import (
"net/http"
"encoding/json"
"fmt"
)
type Payload struct{
Stuff Data}

type Data struct{
Fruit Fruits
Veggies Vegetables}
type Fruits map[string]int
type Vegetables map[string]int

func handler (w http.ResponseWriter, r *http.Request){
response, err:=getJsonResponse()
if err!=nil{
panic(err)
}
fmt.Fprintf(w, string (response))
}
func main(){
http.HandleFunc("/", handler)
http.ListenAndServe("localhost:1339",nil)
}
func getJsonResponse()([]byte, error){
fruits:=make(map[string]int)
fruits["Apples"]=2
fruits["Oranges"]=222
fruits["Pear"]=22
vegetables:=make(map[string]int)
vegetables["potato"]=21
vegetables["tomato"]=32
vegetables["carrot"]=89
d := Data{fruits,vegetables}
p := Payload{d}
return json.MarshalIndent(p,"","  ")
}




