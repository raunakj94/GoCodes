package main

import "net/http"
func response (w http.ResponseWriter,r *http.Request){
w.Write([]byte("Hello Web Server running"))
}
func main(){
http.HandleFunc("/",response)
http.ListenAndServe(":3000",nil)
}

