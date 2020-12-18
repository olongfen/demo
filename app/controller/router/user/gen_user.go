package router

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/demo/app/controller/api/user"
	"github.com/olongfen/demo/app/controller/common"
)

func initUser(r *gin.RouterGroup) {
	c := &api_user.CtrlUser{}
	user := r.Group("user")
	user.GET("get", c.GetOne)
	user.GET("list", c.GetList)
	user.POST("add", c.AddOne)
	user.POST("addList", c.AddList)
	user.PUT("edit", c.Edit)
	user.DELETE("delete", c.DeleteOne)
	user.DELETE("deleteList", c.DeleteList)
}

func init() {
	ctrl_common.RouterGroupFunctions = append(ctrl_common.RouterGroupFunctions, initUser)
}
