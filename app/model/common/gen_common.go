package model_common

import (
	"github.com/olongfen/contrib/log"
	"gorm.io/gorm"
)

type FieldData struct {
	Value  interface{} `json:"value" form:"value"`
	Symbol string      `json:"symbol" form:"symbol"`
}

var (
	ModelLog *log.Logger
	DB       *gorm.DB
	Tables   []interface{}
)

func GetDB(dbs ...*gorm.DB) (ret *gorm.DB) {
	if len(dbs) > 0 {
		return dbs[0]
	}
	return DB
}
