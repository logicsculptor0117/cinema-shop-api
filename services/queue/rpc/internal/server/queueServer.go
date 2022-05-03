// Code generated by goctl. DO NOT EDIT!
// Source: queue.proto

package server

import (
	"context"

	"cinema-shop/services/queue/rpc/internal/logic"
	"cinema-shop/services/queue/rpc/internal/svc"
	"cinema-shop/services/queue/rpc/pb/queue"
)

type QueueServer struct {
	svcCtx *svc.ServiceContext
	queue.UnimplementedQueueServer
}

func NewQueueServer(svcCtx *svc.ServiceContext) *QueueServer {
	return &QueueServer{
		svcCtx: svcCtx,
	}
}

// 创建订单队列
func (s *QueueServer) OrderQueue(ctx context.Context, in *queue.OrderCreateRequest) (*queue.OrderCreateResponse, error) {
	l := logic.NewOrderQueueLogic(ctx, s.svcCtx)
	return l.OrderQueue(in)
}
