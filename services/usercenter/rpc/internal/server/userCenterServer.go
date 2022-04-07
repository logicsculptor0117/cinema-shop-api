// Code generated by goctl. DO NOT EDIT!
// Source: usercenter.proto

package server

import (
	"context"

	"cinema-shop/services/usercenter/rpc/internal/logic"
	"cinema-shop/services/usercenter/rpc/internal/svc"
	"cinema-shop/services/usercenter/rpc/pb/usercenter"
)

type UserCenterServer struct {
	svcCtx *svc.ServiceContext
	usercenter.UnimplementedUserCenterServer
}

func NewUserCenterServer(svcCtx *svc.ServiceContext) *UserCenterServer {
	return &UserCenterServer{
		svcCtx: svcCtx,
	}
}

//  登录验证
func (s *UserCenterServer) GetAccessToken(ctx context.Context, in *usercenter.GetAccessTokenRequest) (*usercenter.GetAccessTokenResponse, error) {
	l := logic.NewGetAccessTokenLogic(ctx, s.svcCtx)
	return l.GetAccessToken(in)
}

//  注册
func (s *UserCenterServer) RegisterUser(ctx context.Context, in *usercenter.RegisterRequest) (*usercenter.RegisterResponse, error) {
	l := logic.NewRegisterUserLogic(ctx, s.svcCtx)
	return l.RegisterUser(in)
}

//  根据ID获取信息
func (s *UserCenterServer) GetUserByID(ctx context.Context, in *usercenter.GetUserByIDRequest) (*usercenter.GetUserByIdResponse, error) {
	l := logic.NewGetUserByIDLogic(ctx, s.svcCtx)
	return l.GetUserByID(in)
}
