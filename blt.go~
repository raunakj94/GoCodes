package main

import (
    "log"
    "fmt"
    "github.com/boltdb/bolt"
"time"
)

func main() {
    // Open the my.db data file in your current directory.
    // It will be created if it doesn't exist.
    db, err := bolt.Open("blog.db", 0600,  &bolt.Options{Timeout: 1 * time.Second})
    if err != nil {
        log.Fatal(err)
    }

db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("posts"))
    if err != nil {
        return err
    }
 d, err := tx.CreateBucketIfNotExists([]byte("post1"))
    if err != nil {
        return err
    }
 f, err := tx.CreateBucketIfNotExists([]byte("post2"))
    if err != nil {
        return err
    }
 g, err := tx.CreateBucketIfNotExists([]byte("post3"))
    if err != nil {
        return err
    }
 h, err := tx.CreateBucketIfNotExists([]byte("post4"))
    if err != nil {
        return err
    }

    return b.Put([]byte("2015-01-01"), []byte("My New Year post"))
 return d.Put([]byte("2015-25-01"), []byte("My Nw Year post"))
 return f.Put([]byte("2015-26-01"), []byte("My ew Year post"))
 return g.Put([]byte("2015-27-01"), []byte("My Ne Year post"))
 return h.Put([]byte("2015-22-01"), []byte("M New Year post"))
})
db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("posts"))
	d:= tx.Bucket([]byte("post1"))
f := tx.Bucket([]byte("post2"))
g := tx.Bucket([]byte("post3"))
h := tx.Bucket([]byte("post4"))

   	v := b.Get([]byte("2015-01-01"))
 	q := d.Get([]byte("2015-25-01"))
	w := f.Get([]byte("2015-26-01"))
	e := g.Get([]byte("2015-27-01"))
	r := h.Get([]byte("2015-22-01"))
    	fmt.Printf("%s", v)
	fmt.Printf("%s", q)
	fmt.Printf("%s", w)
	fmt.Printf("%s", e)
	fmt.Printf("%s", r)

    return nil
})
    defer db.Close()
}

