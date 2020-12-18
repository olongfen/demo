package router

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/demo/app/controller/api/region"
	"github.com/olongfen/demo/app/controller/common"
)

func initRegion(r *gin.RouterGroup) {
	c := &api_region.CtrlRegion{}
	region := r.Group("region")
	region.GET("get", c.GetOne)
	region.GET("list", c.GetList)
	region.POST("add", c.AddOne)
	region.POST("addList", c.AddList)
	region.PUT("edit", c.Edit)
	region.DELETE("delete", c.DeleteOne)
	region.DELETE("deleteList", c.DeleteList)
}

func init() {
	ctrl_common.RouterGroupFunctions = append(ctrl_common.RouterGroupFunctions, initRegion)
}
