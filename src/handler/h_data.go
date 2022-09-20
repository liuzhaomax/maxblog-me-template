package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maxblog-me-template/internal/core"
	"maxblog-me-template/src/service"
	"maxblog-me-template/src/utils"
	"net/http"
)

var DataSet = wire.NewSet(wire.Struct(new(HData), "*"))

type HData struct {
	BData *service.BData
	IRes  core.IResponse
}

func (hData *HData) GetDataById(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := utils.Str2Uint32(idRaw)
	if err != nil {
		hData.IRes.ResFail(c, http.StatusBadRequest, core.FormatError(998, err))
		return
	}
	dataRes, err := hData.BData.GetDataById(c, id)
	if err != nil {
		hData.IRes.ResFail(c, http.StatusInternalServerError, core.FormatError(997, err))
		return
	}
	hData.IRes.ResSuccess(c, gin.H{
		"Hello Data": dataRes.Mobile, // TODO 根据openapi修改
	})
}
