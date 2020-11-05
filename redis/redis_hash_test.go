package redis

import (
	"fmt"
	"testing"
)

func Test_hash_01(t *testing.T)  {
	//hash-------------------------------------------
	hashKey := "userkey_1"
	rdb.HSet(hashKey, "name", "叶子")
	rdb.HSet(hashKey, "age", 18)

	//get hash
	hashGet, _ := rdb.HGet(hashKey, "name").Result()
	fmt.Println("HGet name", hashGet)

	//获取所有hash 返回map
	hashGetAll, _ := rdb.HGetAll(hashKey).Result()
	fmt.Println("HGetAll", hashGetAll)
}