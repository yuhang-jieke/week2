package server

import (
	"context"

	"github.com/yuhang-jieke/week2/wei/user-server/basic/config"
	__ "github.com/yuhang-jieke/week2/wei/user-server/handler/proto"
	"github.com/yuhang-jieke/week2/wei/user-server/model"
)

type Server struct {
	__.UnimplementedUserServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Goods(_ context.Context, in *__.GoodsReq) (*__.GoodsResp, error) {
	goods := model.Goodss{
		Name:  in.Name,
		Price: float64(in.Price),
		Num:   int(in.Num),
		Title: in.Total,
	}
	if err := goods.Create(config.DB); err != nil {
		panic("添加失败")
	}
	return &__.GoodsResp{
		Greet: "添加成功",
	}, nil
}
