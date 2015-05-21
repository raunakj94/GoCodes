package main
import "net/http"
 func response(w http.ResponseWriter,r *http.Request){
w.Write([]byte("runing web server"))
}
func main(){
http.HandleFunc("/",response)
http.ListenAndServe(":3040",nil)
}

type Resource interface{
Get(url url.values) (int , interface{})
Put(url url.values)(int , interface{})
Post(url url.values)(int , interface{})
Delete(url url.values)(int , interface{})
}
type url.values map[string][]string

