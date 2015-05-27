package main
import(
"fmt"
"log"
"github.com/boltdb/bolt"
)
func main(){
db ,err:=bolt.Open("fl.db",0600,nil)
if err!=nil{
  log.Fatal(err)
}
db.Update(func(tx *bolt.Tx) error{
b,err:=tx.CreateBucket([]byte("Total Employee"))
if err!=nil{
return err
}
b.Put([]byte("tech"),[]byte("134"))
b.Put([]byte("non tech"),[]byte("16"))
c:=b.Cursor()
for k,v:=c.First();k !=nil;k, v=c.Next(){
fmt.Printf("the key is %s valueis %s\n",k,v)
}
return nil
})
}
