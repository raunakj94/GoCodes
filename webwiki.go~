package main
import (
"fmt"
"io/ioutil"
"net/http"
)

type Page struct{
Title string
Body []byte
}
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
    filename := title + ".txt"
    body, err:= ioutil.ReadFile(filename)
if err != nil {
        return  nil
    }
    return &Page{Title: title, Body: body}
}
func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2:= loadPage("TestPage")
    fmt.Println(string(p2.Body))
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}



func viewHandler(w http.ResponseWriter, r *http.Request){
title:=r.URL.Path([len("/view/"):])
p:=loadPage(title)
fmt.FPrint(w,<h1>%s</h1><div>%s</div>,title,body)
}

