package main

import (
    "encoding/binary"
    "encoding/json"
    "fmt"
    "log"
    "time"

    "github.com/boltdb/bolt"
)

type Post struct {
    Created time.Time
    Title   string
    Content string
}

type User struct {
    ID   int
    Name string
}

func main() {

    db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    //*
        db.Update(func(tx *bolt.Tx) error {
            b, err := tx.CreateBucketIfNotExists([]byte("posts"))
            if err != nil {
                return err
            }
            return b.Put([]byte("2015-01-01"), []byte("My New Year post"))
        })
    //*/

    db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("posts"))
        v := b.Get([]byte("2015-01-01"))
        fmt.Printf("%s\n", v)

        if err := b.ForEach(func(k, v []byte) error {
            fmt.Printf("A %s is %s.\n", k, v)
            return nil
        }); err != nil {
            return err
        }

        return nil
    })

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

    db.Update(func(tx *bolt.Tx) error {
        b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
        err = b.Put([]byte("answer"), []byte("42"))
        return err
    })

    db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("MyBucket"))
        v := b.Get([]byte("answer"))
        fmt.Printf("%s\n", v)

        return nil
    })

    db.Update(func(tx *bolt.Tx) error {
        u := User{}
        // Retrieve the users bucket.
        // This should be created when the DB is first opened.
        //b := tx.Bucket([]byte("users"))
        b, err := tx.CreateBucketIfNotExists([]byte("users"))

        // Generate ID for the user.
        // This returns an error only if the Tx is closed or not writeable.
        // That can't happen in an Update() call so I ignore the error check.
        id, _ := b.NextSequence()
        u.ID = int(id)

        // Marshal user data into bytes.
        buf, err := json.Marshal(u)
        if err != nil {
            return err
        }

        // Persist bytes to users bucket.
        return b.Put(itob(u.ID), buf)
    })

    // Delete the key in a different write transaction.
    if err := db.Update(func(tx *bolt.Tx) error {
        return tx.Bucket([]byte("MyBucket")).Delete([]byte("answer"))
    }); err != nil {
        log.Fatal(err)
    }

    db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("users"))

        if err := b.ForEach(func(k, v []byte) error {
            fmt.Printf("A %s is %s.\n", k, v)
            return nil
        }); err != nil {
            return err
        }

        return nil
    })

}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(v))
    return b
}