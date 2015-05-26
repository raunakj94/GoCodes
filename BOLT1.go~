package main
import (
"fmt"
"github.com/boltdb/bolt"
"log"
)
func main(){
db,err:=bolt.Open("Bolt1.db",0644,nil)
if err!=nil{
log.Fatal(err)
}
db.Update(func(tx *bolt.Tx) error{
k,err:=tx.CreateBucketIfNotExists([]byte("fission"))
if err!=nil{
return err
}
k.Put([]byte("staff"),[]byte("150"))
k.put([]byte("Size"),[]byte("180"))
})
db.View(func(tx *bolt.Tx) error{
j := tx.Bucket([]byte("fission"))
c:=j.Cursor()
for t, y:=j.First();t!=nil;t,y =j.Next(){
     fmt.Printf("IN %s number of workers are %s.\n",t,y)
}
})
}
      
