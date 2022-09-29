package conf

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"maxblog-me-template/internal/api"
	"maxblog-me-template/internal/middleware/interceptor"
)

func init() {
	gin.DefaultWriter = colorable.NewColorableStdout()
	gin.ForceConsoleColor()
}

func InitGinEngine(iHandler api.IHandler, itcpt *interceptor.Interceptor) *gin.Engine {
	gin.SetMode(cfg.RunMode) // debug, test, release
	app := gin.Default()
	app.Use(LoggerToFile())
	iHandler.Register(app, itcpt)
	return app
}
