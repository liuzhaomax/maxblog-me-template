package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maxblog-me-template/internal/core"
	"maxblog-me-template/src/schema"
	"maxblog-me-template/src/service"
	"net/http"
)

var DataSet = wire.NewSet(wire.Struct(new(HData), "*"))

type HData struct {
	BData *service.BData
	IRes  core.IResponse
}

func (hData *HData) GetDataById(c *gin.Context) {
	var dataReq schema.DataReq
	if err := c.ShouldBind(&dataReq); err != nil {
		hData.IRes.ResFail(c, http.StatusBadRequest, core.NewError(998, err))
	}
	dataRes, err := hData.BData.GetDataById(c, &dataReq)
	if err != nil {
		hData.IRes.ResFail(c, http.StatusBadRequest, core.NewError(998, err))
	}
	hData.IRes.ResSuccess(c, gin.H{
		"Hello Data": dataRes.Mobile, // TODO 根据openapi修改
	})
}
