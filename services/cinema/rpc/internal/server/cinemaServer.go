// Code generated by goctl. DO NOT EDIT!
// Source: cinema.proto

package server

import (
	"context"

	"cinema-shop/services/cinema/rpc/internal/logic"
	"cinema-shop/services/cinema/rpc/internal/svc"
	"cinema-shop/services/cinema/rpc/pb/cinema"
)

type CinemaServer struct {
	svcCtx *svc.ServiceContext
	cinema.UnimplementedCinemaServer
}

func NewCinemaServer(svcCtx *svc.ServiceContext) *CinemaServer {
	return &CinemaServer{
		svcCtx: svcCtx,
	}
}

// 影片列表
func (s *CinemaServer) FilmList(ctx context.Context, in *cinema.FilmListRequest) (*cinema.FilmListResponse, error) {
	l := logic.NewFilmListLogic(ctx, s.svcCtx)
	return l.FilmList(in)
}

// 影片详情
func (s *CinemaServer) FilmDetail(ctx context.Context, in *cinema.FilmDatailRequest) (*cinema.FilmDetailResponse, error) {
	l := logic.NewFilmDetailLogic(ctx, s.svcCtx)
	return l.FilmDetail(in)
}

// 根据地理位置获取影院信息
func (s *CinemaServer) CinemaInfo(ctx context.Context, in *cinema.CinemaInfoRequest) (*cinema.CinemaInfoResponse, error) {
	l := logic.NewCinemaInfoLogic(ctx, s.svcCtx)
	return l.CinemaInfo(in)
}

// 根据影片和日期和影院ID获取排片信息
func (s *CinemaServer) ScreenCinemaInfo(ctx context.Context, in *cinema.ScreenCinemaInfoRequest) (*cinema.ScreenCinemaInfoResponse, error) {
	l := logic.NewScreenCinemaInfoLogic(ctx, s.svcCtx)
	return l.ScreenCinemaInfo(in)
}
