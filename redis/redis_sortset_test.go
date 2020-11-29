package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

func Test_zset_01(t *testing.T) {
	//zset------------------------------------------------
	zsetKey := "go2zset"
	ranking := []redis.Z{
		redis.Z{Score: 100.0, Member: "钟南山"},
		redis.Z{Score: 80.0, Member: "林医生"},
		redis.Z{Score: 70.0, Member: "王医生"},
		redis.Z{Score: 75.0, Member: "张医生"},
		redis.Z{Score: 59.0, Member: "叶医生"},
	}
	rdb.ZAdd(zsetKey, ranking...)
	//golang+5分
	newScore, _ := rdb.ZIncrBy(zsetKey, 5.0, "钟南山").Result()
	fmt.Println("钟南山加5分后的最新分数", newScore)
	//取zset里的前2名热度的医生
	zsetList2, _ := rdb.ZRevRangeWithScores(zsetKey, 0, 1).Result()
	fmt.Println("zset前2名热度的医生", zsetList2)
}
