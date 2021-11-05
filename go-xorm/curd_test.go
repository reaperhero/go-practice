package go_xorm__test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
	"xorm.io/xorm"
)

type User struct {
	Id       int       `xorm:"not null pk autoincr INT(11)"`
	Username string    `xorm:"not null VARCHAR(32)"`
	Birthday time.Time `xorm:"DATE"`
	Sex      string    `xorm:"CHAR(1)"`
	Address  string    `xorm:"VARCHAR(256)"`
}

func Test_Xasd(t *testing.T) {

	//创建orm引擎
	engine, err := xorm.NewEngine("mysql", "root:Lzslov123!@/ggg?charset=utf8")

	if err != nil {
		fmt.Println(err)
		return
	}
	//连接测试
	if err := engine.Ping(); err != nil {
		fmt.Println(err)
		return
	}

	//日志打印SQL
	engine.ShowSQL(true)

	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	engine.SetMaxOpenConns(5)


	err = engine.Sync2(new(User))

	//增
	user := new(User)
	user.Username="tyming"
	affected,err := engine.Insert(user)
	fmt.Println(affected)

	//删
	//user := new(User)
	//user.Username="tyming"
	//affected_delete,err := engine.Delete(user)
	//fmt.Println(affected_delete)

	//改
	//user := new(User)
	//user.Username="tyming"
	//affected_update,err := engine.Id(1).Update(user)
	//fmt.Println(affected_update)

	//查
	user = new(User)
	//result,err := engine.Id(1).Get(user)
	result, err := engine.Where("id=?", 1).Get(user)
	fmt.Println(result)

}
