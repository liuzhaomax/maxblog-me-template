package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	srcHandler "maxblog-me-template/src/handler"
)

var APISet = wire.NewSet(wire.Struct(new(Handler), "*"), wire.Bind(new(IHandler), new(*Handler)))

type Handler struct {
	HandlerData *srcHandler.HData
}

type IHandler interface {
	Register(app *gin.Engine)
}

func (handler *Handler) Register(app *gin.Engine) {
	handler.RegisterRouter(app)
}
