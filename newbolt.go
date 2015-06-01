package main
import "fmt"
import "github.com/boltdb/bolt"
import "log"

func main(){
db,err:=bolt.Open("abc.db",0600,nil)
if err!=nil{
return log.Fatal(err)
}
db.Update(func (tx *bolt.Tx) error{
d,err=tx.CreateBucket([]byte("subline id"))


