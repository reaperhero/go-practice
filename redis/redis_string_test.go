package redis

import (
	"encoding/json"
	"fmt"
	//redis "gopkg.in/redis.v4"
	"testing"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "Jsx4ujds2P8veOCgz",
		DB:       0,
		PoolSize: 20,
	})
}

// string
func Test_string_01(t *testing.T) {
	// set
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
	// get
	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	// get none
	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

// inc

// struct
func Test_struct_01(t *testing.T) {
	type Doctor struct {
		Id   int
		Name string
		Age  int
		Time time.Time
	}
	doctor := Doctor{1, "钟南山", 83, time.Now()}
	doctorJson, _ := json.Marshal(doctor)
	rdb.Set("doctor2", doctorJson, time.Hour)

	//读取结构
	doctorResult, _ := rdb.Get("doctor2").Result()
	var doctor2 Doctor
	//反序列化
	json.Unmarshal([]byte(doctorResult), &doctor2)
	fmt.Println("doctor2", doctor2)
}
