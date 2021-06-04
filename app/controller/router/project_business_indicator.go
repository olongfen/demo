package router

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/demo/app/controller/apis"
)

func initProjectBusinessIndicator(r *gin.RouterGroup) {
	c := &apis.CtlProjectBusinessIndicator{}
	projectBusinessIndicator := r.Group("project_business_indicators")
	projectBusinessIndicator.GET(":id", c.Get)
	projectBusinessIndicator.GET("", c.GetBatch)
	projectBusinessIndicator.POST("", c.Create)
	projectBusinessIndicator.POST("batch", c.AddBatch)
	projectBusinessIndicator.PUT(":id", c.Update)
	// projectBusinessIndicator.DELETE(":id", c.Delete) // default close
	// projectBusinessIndicator.DELETE("batch", c.DeleteBatch) // default close
}
func init() {
	apis.RouterGroupFunctions = append(apis.RouterGroupFunctions, initProjectBusinessIndicator)
}
