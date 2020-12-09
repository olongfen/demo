package router

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/demo/app/controller/api/admin"
)

func initAdmin(r *gin.RouterGroup) {
	c := &api_admin.CtrlAdmin{}
	admin := r.Group("admin")
	admin.GET("get", c.GetOne)
	admin.GET("list", c.GetList)
	admin.POST("add", c.AddOne)
	admin.POST("addList", c.AddList)
	admin.PUT("edit", c.Edit)
	admin.DELETE("delete", c.DeleteOne)
	admin.DELETE("deleteList", c.DeleteList)
}
