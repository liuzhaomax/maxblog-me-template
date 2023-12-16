package router

import (
	"github.com/gin-gonic/gin"
	"maxblog-me-template/src/handler"
)

func RegisterRouter(handler *handler.HData, app *gin.Engine) {
	//itcpt := &interceptor.Interceptor{}
	routerData := app.Group("")
	{
		routerData.GET("/:id", handler.GetDataById)
	}
}
