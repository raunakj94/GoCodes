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

  key, err := NewKey("test", "aerospike", "key")
  panicOnError(err)

  bins := BinMap{
    "bin1": 42,
    "bin2": 27,
    "bin3": 63,
  }

 
  err = client.Put(nil, key, bins)
  panicOnError(err)

  rec, err := client.Get(nil, key)
  panicOnError(err)

  fmt.Printf("%#v\n", *rec)

}
