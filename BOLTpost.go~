package main

import (
    "log"
    "fmt"
    "github.com/boltdb/bolt"
"time"
 "encoding/json"
)
func main(){
type Post struct {
    Created time.Time
    Title   string
    Content string
}
db, err :=bolt.Open("post.db",0600,nil)
	if err!=nil{
	log.Fatal(err)
}
defer db.Close()

post := &Post{
   Created: time.Now(),
   Title:   "My first post",
   Content: "Hello, this is my first post.",
}

db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("posts"))
    if err != nil {
        return err
    }
    encoded, err := json.Marshal(post)
    if err != nil {
        return err
    }
    return b.Put([]byte(post.Created.Format(time.RFC3339)), encoded)
})
err = db.View(func(tx *bolt.Tx) error {
      j:=tx.Bucket([]byte("posts"))
     o, err := decode(b.Get(k))
        if err != nil {
            return err
        }
    fmt.Println(o)
return nil
})
func decode(data []byte) (*Post, error) {
    var p *Post
    err := json.Unmarshal(data, &p)
    if err != nil {
        return err
    }
    return p
}
}
