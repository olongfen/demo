package main

import "gorm.io/gorm"

//go:generate gengo  -i ./gen.go -output . -m github.com/olongfen/demo
type User struct {
	gorm.Model
	Name  string
	Age   int
	Class string
}

type Admin struct {
	gorm.Model
	Name  string
	Age   int
	Class string
}
