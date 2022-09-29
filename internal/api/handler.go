package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maxblog-me-template/internal/middleware/interceptor"
	srcHandler "maxblog-me-template/src/handler"
)

var APISet = wire.NewSet(wire.Struct(new(Handler), "*"), wire.Bind(new(IHandler), new(*Handler)))

type Handler struct {
	HandlerData *srcHandler.HData
}

type IHandler interface {
	Register(app *gin.Engine, itcpt *interceptor.Interceptor)
}

func (handler *Handler) Register(app *gin.Engine, itcpt *interceptor.Interceptor) {
	handler.RegisterRouter(app, itcpt)
}
