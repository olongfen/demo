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
)
