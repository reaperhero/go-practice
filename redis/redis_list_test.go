package redis

import (
	"fmt"
	"testing"
	"time"
)

func Test_list_01(t *testing.T) {
	listKey := "go2list"
	rdb.RPush(listKey, "str1", 10, "str2", 15, "str3", 20).Err()

	//lpop 取出并移除左边第一个元素
	first, _ := rdb.LPop(listKey).Result()
	fmt.Printf("列表第一个元素 key:%v value:%v \n", first[0], first[1])

	//Blpop 取出并移除左边第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	first2, _ := rdb.BLPop(time.Second*60, listKey).Result()
	fmt.Printf("列表第一个元素 key:%v value:%v \n", first2[0], first2[1])

	//数据长度
	listLen, _ := rdb.LLen(listKey).Result()
	fmt.Println("list length", listLen)

	//获取列表
	listGet, _ := rdb.LRange(listKey, 1, 2).Result()
	fmt.Println("索引1-2的2个元素元素", listGet)
}
