package router

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/olongfen/contrib/log"
	"github.com/olongfen/demo/app/controller/apis"
	"github.com/olongfen/demo/app/controller/middleware"
	"github.com/olongfen/demo/app/setting"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Engine 初始化路由
var (
	Engine = gin.Default()
)

// Init 初始化路由模块
func Init() {
	if !setting.DevEnv {
		gin.SetMode(gin.ReleaseMode)
		Engine.Use(gin.Logger())
		// 创建记录日志的文件
		f, _ := rotatelogs.New(
			setting.Global.FilePath.LogDir + "/router" + setting.Global.FilePath.LogPatent + ".log",
		)
		gin.DefaultWriter = io.MultiWriter(f)
	}

	// 添加中间件
	Engine.Use(middleware.CORS())
	Engine.Use(middleware.GinLogFormatter())
	Engine.Use(gin.Recovery())
	// 没有路由请求
	Engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": " 404 " + http.StatusText(http.StatusNotFound),
		})
	})
	// TODO 路由
	{
		var api = Engine.Group("api/v1")
		api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.Use(middleware.Common())

		// 测试连接
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"ping": "pong >>>>>>> update"})
		})
		for _, v := range apis.RouterGroupFunctions {
			v(api)
		}

	}
	log.Infoln("router init success !")
}
