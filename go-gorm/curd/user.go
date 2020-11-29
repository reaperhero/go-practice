package curd

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      *int
	Birthday time.Time
	//Email        string  `gorm:"type:varchar(100);unique_index"`
	Role string `gorm:"size:255"` //设置字段的大小为255个字节
	//MemberNumber string `gorm:"unique;not null"` // 设置 memberNumber 字段唯一且不为空
	//Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 Num字段自增
	//Address      string  `gorm:"index:addr"`      // 给Address 创建一个名字是  `addr`的索引
	//IgnoreMe     int     `gorm:"-"`               //忽略这个字段
}

// 设置 `User` 的表名为 `profiles`
func (user *User) TableName() string {
	return "user"
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("id", uuid.New())
	return nil
}

type Email struct {
	gorm.Model
	Name string
}
