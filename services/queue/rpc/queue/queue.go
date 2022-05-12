// Code generated by goctl. DO NOT EDIT!
// Source: queue.proto

package queue

import (
	"context"

	"cinema-shop/services/queue/rpc/pb/queue"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	OrderCreateRequest  = queue.OrderCreateRequest
	OrderCreateResponse = queue.OrderCreateResponse
	OrderDelayRequest   = queue.OrderDelayRequest
	OrderDelayResponse  = queue.OrderDelayResponse

	Queue interface {
		// 创建订单队列
		OrderQueue(ctx context.Context, in *OrderCreateRequest, opts ...grpc.CallOption) (*OrderCreateResponse, error)
		// 自动取消未支付订单队列
		OrderDelay(ctx context.Context, in *OrderDelayRequest, opts ...grpc.CallOption) (*OrderDelayResponse, error)
	}

	defaultQueue struct {
		cli zrpc.Client
	}
)

func NewQueue(cli zrpc.Client) Queue {
	return &defaultQueue{
		cli: cli,
	}
}

// 创建订单队列
func (m *defaultQueue) OrderQueue(ctx context.Context, in *OrderCreateRequest, opts ...grpc.CallOption) (*OrderCreateResponse, error) {
	client := queue.NewQueueClient(m.cli.Conn())
	return client.OrderQueue(ctx, in, opts...)
}

// 自动取消未支付订单队列
func (m *defaultQueue) OrderDelay(ctx context.Context, in *OrderDelayRequest, opts ...grpc.CallOption) (*OrderDelayResponse, error) {
	client := queue.NewQueueClient(m.cli.Conn())
	return client.OrderDelay(ctx, in, opts...)
}