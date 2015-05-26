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

    b.Put([]byte("BioTech"), []byte("67"))
    b.Put([]byte("cat"), []byte("lame"))
    b.Put([]byte("liger"), []byte("awesome"))
    c := b.Cursor()
    for k, v := c.First(); k != nil; k, v = c.Next() {
        fmt.Printf("A %s is %s.\n", k, v)
    }
    m.Put([]byte("Finance"),[]byte("22"))
 m.Put([]byte("Marketing"),[]byte("222"))
 m.Put([]byte("Eco"),[]byte("224"))
j:=m.Cursor()
for t, y:=j.First();t!=nil;t,y =j.Next(){
     fmt.Printf("IN %s number of students are %s.\n",t,y)
}

    return nil
})
db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("bTech"))
    if err != nil {
        return err
    }
  return b.Put([]byte("CSE"), []byte("68"))

})
db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("bTech"))
    if err != nil {
        return err
    }
  return b.Put([]byte("IT"), []byte("69"))

})
err = db.View(func(tx *bolt.Tx) error {
      j:=tx.Bucket([]byte("bTech"))
      c := j.Cursor()
    	for k, n := c.First(); k != nil; k, n = c.Next() {
        	fmt.Printf("key=%s, value=%s\n", k, n)
    		}
return nil
})
}
