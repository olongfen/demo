package srv_common

import (
	"github.com/olongfen/contrib/log"
	"github.com/olongfen/demo/app/setting"
	"gorm.io/gorm"
)

type FieldData struct {
	Value  interface{} `json:"value" form:"value"`
	Symbol string      `json:"symbol" form:"symbol"`
}

var (
	ServiceLog = log.NewLogFile(log.ParamLog{Path: setting.Global.FilePath.LogDir + "/" + "service", Stdout: !setting.DevEnv, P: setting.Global.FilePath.LogPatent})
	DB         *gorm.DB
)
