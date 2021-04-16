package main

import (
	"context"
	"fmt"

	"github.com/golang/groupcache"
)

type SlowDB struct {
	data map[string]string
}

func (db *SlowDB) Get(key string) string {
	fmt.Println("get " + key + " from db")
	return db.data[key]
}

func (db *SlowDB) Set(key string, value string) {
	db.data[key] = value
}

func NewSlowDB() *SlowDB {
	ndb := new(SlowDB)
	ndb.data = make(map[string]string)
	return ndb
}

func main() {

	db := NewSlowDB()

	db.Set("foo", "bar")
	db.Set("one", "two")

	var stringcache = groupcache.NewGroup("SlowDBCache", 64<<20, groupcache.GetterFunc(
		func(ctx context.Context, key string, dest groupcache.Sink) error {
			result := db.Get(key) // 数据库模拟操作
			dest.SetBytes([]byte(result))
			return nil
		}))

	var data []byte

	err := stringcache.Get(nil, "foo", groupcache.AllocatingByteSliceSink(&data))  // get foo from db
	stringcache.Get(nil, "foo", groupcache.AllocatingByteSliceSink(&data))         // no print
	err2 := stringcache.Get(nil, "one", groupcache.AllocatingByteSliceSink(&data)) // get one from db

	db.Set("foo", "bar2")
	err3 := stringcache.Get(nil, "foo", groupcache.AllocatingByteSliceSink(&data)) // data was bar

	if err != nil {
		fmt.Println("error")
	}

	if err2 != nil {
		fmt.Println("error2")
	}

	if err3 != nil {
		fmt.Println("error3")
	}

	fmt.Printf("data was %s\n", data)

}
