package main

import (
    "log"
    "fmt"
    "github.com/boltdb/bolt"
    "time"
)
func main(){
db, err :=bolt.Open("dfm.db",0600,&bolt.Options{Timeout: 1 * time.Second})
	if err!=nil{
	log.Fatal(err)
}
defer db.Close()

db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("AdServer"))
    if err != nil {
        return err
    }
    b.Put([]byte("abc"), []byte("67"))
    b.Put([]byte("abc1"), []byte("97.8"))
    b.Put([]byte("abc2"), []byte("85.89"))
    c := b.Cursor()
    for k, v := c.First(); k != nil; k, v = c.Next() {
        fmt.Printf("A %s is %s.\n", k, v)
    }
 return nil
})

db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("bTech"))
    if err != nil {
        return err
    }
  return b.Put([]byte("abc3"), []byte("68"))

})
err = db.View(func(tx *bolt.Tx) error {
      j:=tx.Bucket([]byte("subline_id"))
      c := j.Cursor()
    	for k, n := c.First(); k != nil; k, n = c.Next() {
        	fmt.Printf("key=%s, value=%s\n", k, n)
    		}
return nil
})
}
