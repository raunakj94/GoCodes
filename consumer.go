package main
import (
"fmt"
"encoding/json"
"io/ioutil"
)
type Payload struct{
Stuff Data}

type Data struct{
Fruit Fruits
Veggies Vegetbles
}
type Fruits map[string]int
type Vegetablesmap[string]int

func main(){

