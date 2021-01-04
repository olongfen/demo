package router

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/demo/app/controller/api/demo"
	"github.com/olongfen/demo/app/controller/common"
)

func initDemo(r *gin.RouterGroup) {
	c := &api_demo.CtrlDemo{}
	demo := r.Group("demo")
	demo.GET("get", c.GetOne)
	demo.GET("list", c.GetList)
	demo.POST("add", c.AddOne)
	demo.POST("addList", c.AddList)
	demo.PUT("edit", c.Edit)
	demo.DELETE("delete", c.DeleteOne)
	demo.DELETE("deleteList", c.DeleteList)
}

func init() {
	ctrl_common.RouterGroupFunctions = append(ctrl_common.RouterGroupFunctions, initDemo)
}
