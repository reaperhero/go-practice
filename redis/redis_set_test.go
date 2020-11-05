package redis

import (
	"fmt"
	"testing"
)

func Test_set_01(t *testing.T)  {
	//set-------------------------------------------- 
	setKey := "go2set"
	rdb.SAdd(setKey, "set1")
	rdb.SAdd(setKey, "set2")
	rdb.SAdd(setKey, "set3")
	rdb.SAdd(setKey, "set4")

	//获取集合的所有成员  
	setList, _ := rdb.SMembers(setKey).Result()
	fmt.Println("GetSet", setList)
	//移除集合里的set1  
	rdb.SRem(setKey, "set1")

	//移除并返回set的一个随机元素  
	setFirst, _ := rdb.SPop(setKey).Result()
	fmt.Println("setFirst", setFirst)
}