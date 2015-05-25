package main

import (
    "log"
    "fmt"
    "github.com/boltdb/bolt"
)
func main(){
db, err :=bolt.Open("nit.db",0600,nil)
	if err!=nil{
	log.Fatal(err)
}
defer db.Close()

db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("bTech"))
    if err != nil {
        return err
    }
m, err:=tx.CreateBucketIfNotExists([]byte("MBA"))
    if err != nil {
	return err
    }

  return b.Put([]byte("BioTech"), []byte("67"))
return b.Put([]byte("CSE"), []byte("68"))
return b.Put([]byte("IT"), []byte("69"))

  return m.Put([]byte("FINance"), []byte("671"))
return m.Put([]byte("marketing"), []byte("681"))
return m.Put([]byte("eco"), []byte("691"))
})
err = db.View(func(tx *bolt.Tx) error {
      j:=tx.Bucket([]byte("bTech"))
      v:=j.Get([]byte("IT"))
      fmt.Printf("%s", v)
return nil
})
}
