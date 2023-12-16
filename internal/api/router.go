package api

import (
	"github.com/gin-gonic/gin"
	mw "maxblog-me-template/internal/middleware"
	srcRouter "maxblog-me-template/src/router"
	"net/http"
)

func (handler *Handler) RegisterRouter(app *gin.Engine) {
	app.NoRoute(handler.GetNoRoute)
	app.Use(mw.Cors())
	app.StaticFS("/static", http.Dir("./static"))
	srcRouter.RegisterRouter(handler.HandlerData, app)
}

func (handler *Handler) GetNoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"res": "404"})
}
