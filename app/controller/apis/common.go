package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/log"
	"github.com/olongfen/demo/app/setting"
)

var (
	ControlLog           = log.NewLogFile(log.ParamLog{Path: setting.Global.FilePath.LogDir + "/" + "controller", Stdout: setting.DevEnv, P: setting.Global.FilePath.LogPatent})
	RouterGroupFunctions []func(group *gin.RouterGroup)
)
