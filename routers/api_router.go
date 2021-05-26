package routers

import (
	"github.com/konger/ckgo/controller/v1/api"

	"github.com/gin-gonic/gin"
)

func RegisterRouterApi(app *gin.RouterGroup) {
	user := api.User{}
	app.POST("/user/register", user.Register)
	//app.POST("/user/login", user.Login)
	app.POST("/user/logout", user.Logout)

}
