package main

import(
	"fmt"
	"github.com/dgraph-io/badger"
	"io/ioutil"
)

func main(){
	opt := badger.DefaultOptions
	dir, _ := ioutil.TempDir("/tmp", "badger")
	opt.Dir = dir
	opt.ValueDir = dir
	kv, _ := badger.NewKV(&opt)

	key := []byte("hello")

	kv.Set(key, []byte("world"))
	fmt.Printf("SET %s world\n", key)

	var item badger.KVItem
	if err := kv.Get(key, &item); err != nil {
		fmt.Printf("Error while getting key: %q", key)
	}
	fmt.Printf("GET %s %s\n", key, item.Value())

	if err := kv.CompareAndSet(key, []byte("venus"), 100); err != nil {
		fmt.Println("CAS counter mismatch")
	} else {
		if err = kv.Get(key, &item); err != nil {
			fmt.Printf("Error while getting key: %q", key)
		}
		fmt.Printf("Set to %s\n", item.Value())
	}
	if err := kv.CompareAndSet(key, []byte("mars"), item.Counter()); err == nil {
		fmt.Println("Set to mars")
	} else {
		fmt.Printf("Unsuccessful write. Got error: %v\n", err)
	}	
}