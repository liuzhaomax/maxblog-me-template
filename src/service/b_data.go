package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"maxblog-me-template/internal/core"
	"maxblog-me-template/src/pb"
	"maxblog-me-template/src/schema"
	"maxblog-me-template/src/utils"
)

var DataSet = wire.NewSet(wire.Struct(new(BData), "*"))

type BData struct{}

func (b *BData) GetDataById(c *gin.Context, id uint32) (*schema.DataRes, error) {
	addr := core.GetDownstreamMaxblogBETemplateAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("[%s 方法失败] %s: %s", utils.GetFuncName(), core.GRPC_Dial_Failed, err.Error())
		return nil, err
	}
	client := pb.NewDataServiceClient(conn)
	pbRes, err := client.GetDataById(context.Background(), &pb.IdRequest{Id: id})
	if err != nil {
		return nil, err
	}
	dataRes := schema.Pb2Res(pbRes)
	return &dataRes, nil
}
