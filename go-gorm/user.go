package go_gorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name         string
	Age          int
	Birthday     time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段的大小为255个字节
	MemberNumber string `gorm:"unique;not null"` // 设置 memberNumber 字段唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 Num字段自增
	Address      string  `gorm:"index:addr"`      // 给Address 创建一个名字是  `addr`的索引
	IgnoreMe     int     `gorm:"-"`               //忽略这个字段
}

// 设置 `User` 的表名为 `profiles`
func (User) TableName() string {
	return "table_user"
}

func init() {
	db, _ = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dbuser?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&User{})
	db.SingularTable(true) // 禁用表名复数形式
}
