// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"cinema-shop/services/user/rpc/internal/logic"
	"cinema-shop/services/user/rpc/internal/svc"
	"cinema-shop/services/user/rpc/pb/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

//  根据邮箱获取用户秘钥和密码
func (s *UserServer) GetUserByEmail(ctx context.Context, in *user.GetUserByEmailRequest) (*user.GetUserByEmailResponse, error) {
	l := logic.NewGetUserByEmailLogic(ctx, s.svcCtx)
	return l.GetUserByEmail(in)
}

//  根据ID获取信息
func (s *UserServer) GetUserByID(ctx context.Context, in *user.GetUserByIDRequest) (*user.GetUserByIdResponse, error) {
	l := logic.NewGetUserByIDLogic(ctx, s.svcCtx)
	return l.GetUserByID(in)
}
