package main
import (
"fmt"
"io/ioutil"
)
type Page struct{
Title string
Body []byte}
func (p *Page) save() error{
filename:=p.Title+".txt"
return ioutil.Writefile(filename,p.Body,0600)
}
func loadpage(title string) *page{
filename:=title+".txt"
body ,err=ioutil.Readfile(filename)
if err!=nil{
         return nil
}
return &page(Title:title,Body:body)
}
func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2:= loadPage("TestPage")
    fmt.Println(string(p2.Body))
}
