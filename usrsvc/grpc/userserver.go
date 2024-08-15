package grpc

import (
	"context"
	protos "proto"
	"usrsvc/dao"

	"github.com/spf13/cast"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	protos.UnimplementedUserServiceServer
}

func (u *UserServer) GetUser(ctx context.Context, in *protos.User) (*protos.User, error) {
	info := dao.UserDao.GetUserById(ctx, in.Id)
	if len(info) == 0 {
		return nil, status.Errorf(404, "user not found")
	}
	return &protos.User{
		Id:   cast.ToInt64(info["id"]),
		Name: cast.ToString(info["name"]),
	}, nil
}
