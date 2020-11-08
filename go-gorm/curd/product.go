package curd

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Price float64
}
