package grpc

import (
	"context"
	protos "proto"
	"skusvc/dao"

	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
)

type SkuServer struct {
	protos.UnimplementedSkuServiceServer
}

func (s *SkuServer) DecreaseStock(ctx context.Context, in *protos.Sku) (*protos.Sku, error) {

	info := dao.SkuDao.GetSkuById(ctx, in.Id)

	if len(info) == 0 {
		return nil, status.Errorf(404, "sku not found")
	}
	desrRes, err := dao.SkuDao.Decr(ctx, in.Id, in.Num)
	if err != nil {
		return nil, status.Errorf(500, "decrease stock failed")
	}
	if affeced, _ := desrRes.RowsAffected(); affeced == 0 {
		return nil, status.Errorf(500, "decrease stock failed")
	}
	return &protos.Sku{
		Id:    cast.ToInt64(info["id"]),
		Num:   cast.ToInt32(info["num"]) - in.Num,
		Name:  cast.ToString(info["name"]),
		Price: cast.ToInt32(info["price"]),
	}, nil
}
