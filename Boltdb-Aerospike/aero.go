package main

import (
  "fmt"

  . "github.com/aerospike/aerospike-client-go"
)

func panicOnError(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  client, err := NewClient("127.0.0.1", 3000)
  panicOnError(err)

  key, err := NewKey("dfm", "adserver", "key")
  panicOnError(err)
  bins := BinMap{
    "abc1": 67,
    "abc2": 97.8,
    "abc3": 85.89,
  }
  err = client.Put(nil, key, bins)
  panicOnError(err)
bin2 := NewBin("abc4", 25)
err=client.Put(nil ,key, bin2)
  rec, err := client.Get(nil, key)
  panicOnError(err)

  fmt.Printf("%#v\n", *rec)
  existed, err := client.Delete(nil, key)
  panicOnError(err)
  fmt.Printf("Record existed before delete? %v\n", existed)
}
