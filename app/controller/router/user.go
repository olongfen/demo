package router

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/demo/app/controller/apis"
)

func initUser(r *gin.RouterGroup) {
	c := &apis.CtlUser{}
	user := r.Group("users")
	user.GET(":id", c.Get)
	user.GET("", c.GetBatch)
	user.POST("", c.Create)
	user.POST("batch", c.AddBatch)
	user.PUT(":id", c.Update)
	// user.DELETE(":id", c.Delete) // default close
	// user.DELETE("batch", c.DeleteBatch) // default close
}
func init() {
	apis.RouterGroupFunctions = append(apis.RouterGroupFunctions, initUser)
}
