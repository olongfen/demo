package main

import "gorm.io/gorm"

//go:generate gengo  -i ./gen.go -output . -m github.com/olongfen/demo
type User struct {
	gorm.Model
	Name  string `grom:"unique_index"`
	Age   int
	Class string
}

type Admin struct {
	gorm.Model
	Name  string
	Age   int
	Class string
}

type Region struct {
	ID        int64  `json:"id" gorm:"primarykey"`
	Name      string `json:"name" gorm:"type:varchar(64)"`            //
	Cname     string `json:"cname" gorm:"type:varchar(64)"`           //
	LowerName string `json:"lowerName" gorm:"type:varchar(64)"`       //
	Longitude string `json:"longitude" gorm:"type:varchar(12);index"` // 经度
	Latitude  string `json:"latitude" gorm:"type:varchar(12);index"`  // 纬度
}
