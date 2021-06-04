package services

import (
	"github.com/olongfen/contrib/log"
	"github.com/olongfen/demo/app/setting"
)

var (
	ServiceLog = log.NewLogFile(log.ParamLog{Path: setting.Global.FilePath.LogDir + "/" + "service", Stdout: setting.DevEnv, P: setting.Global.FilePath.LogPatent})
)
