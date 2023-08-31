package main

import (
	"context"
	"fmt"
	"time"

	// 00.引入数据库驱动
	_ "gitee.com/chunanyong/dm"
	"gitee.com/chunanyong/zorm"
	log "github.com/sirupsen/logrus"
)

var DbDao *zorm.DBDao

// 01.初始化DBDao
func main() {
	dbConfig := zorm.DataSourceConfig{
		//连接数据库DSN
		DSN: "dm://SYSDBA:SYSDBA001@172.16.104.165:5236", //DM
		//DSN: "root:root@tcp(127.0.0.1:3306)/zorm_t?charset=utf8&parseTime=true&loc=Local", //MySQL
		//数据库类型
		DriverName: "dm",
		//DBType:     "dm",
		//sql.Open(DriverName,DSN) DriverName就是驱动的sql.Open第一个字符串参数,根据驱动实际情况获取
		//DriverName: "mysql",
		Dialect: "dm",
		//SlowSQLMillis 慢sql的时间阈值,单位毫秒.小于0是禁用SQL语句输出;等于0是只输出SQL语句,不计算执行时间;大于0是计算SQL执行时间,并且>=SlowSQLMillis值
		SlowSQLMillis: 0,
		//最大连接数 默认50
		MaxOpenConns: 0,
		//最大空闲数 默认50
		MaxIdleConns: 0,
		//连接存活秒时间. 默认600
		ConnMaxLifetimeSecond: 0,
		//事务隔离级别的默认配置,默认为nil
		DefaultTxOptions: nil,
	}

	var err error
	DbDao, err = zorm.NewDBDao(&dbConfig)
	if err != nil {
		log.Fatalf("数据库连接异常 %v", err)
	}
	log.Println("数据库连接成功")
	c := context.Background()
	ctx, err := DbDao.BindContextDBConnection(c)
	if err != nil {
		panic(err)
	}

	entityMap, err := zorm.InsertEntityMap(ctx, &AgentList{
		Id:             "",
		SidecarId:      "",
		Type:           0,
		Name:           "",
		Version:        "",
		IsUninstalled:  0,
		DeployDate:     time.Time{},
		AutoDeployment: 0,
		LastUpdateDate: time.Time{},
		AutoUpdated:    0,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(entityMap)
}
