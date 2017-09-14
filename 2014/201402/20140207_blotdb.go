package main

import (
    "log"

    "github.com/boltdb/bolt"
    "bytes"
    "fmt"
)

type Cache struct{
    db *bolt.DB;
    requestsBucket []byte;
}

func main() {
    db, err := bolt.Open("my.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    var cache Cache;
    cache.Bind(db);
    cache.Set([]byte("first"), []byte("ok"));

    get, _:= cache.Get([]byte("first"));
    fmt.Println("get: ", string(get));
    return;
}

func (c *Cache) Bind(Db *bolt.DB){
    c.db = Db;
    c.requestsBucket = []byte("test");
}

func (c *Cache) Set(key, value []byte) error {
    err := c.db.Update(func(tx *bolt.Tx) error {
        bucket, err := tx.CreateBucketIfNotExists(c.requestsBucket)
        if err != nil {
            return err
        }
        err = bucket.Put(key, value)
        if err != nil {
            return err
        }
        return nil
    })
 
    return err
}
 
func (c *Cache) Get(key []byte) (value []byte, err error) {
   err = c.db.View(func(tx *bolt.Tx) error {
      bucket := tx.Bucket(c.requestsBucket)
      if bucket == nil {
         return fmt.Errorf("Bucket %q not found!", c.requestsBucket)
      }
 
      var buffer bytes.Buffer
      buffer.Write(bucket.Get(key))
 
      value = buffer.Bytes()
      return nil
   })
 
   return
}