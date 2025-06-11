package grpc

import (
	"context"
	paymentpb "order-service/internal/adapters/grpc/gen/go"
	"order-service/internal/core/ports"

	"google.golang.org/grpc"
)

type PaymentGrpcClient struct {
	client paymentpb.PaymentServiceClient
}

func NewPaymentGrpcClient(conn *grpc.ClientConn) ports.PaymentService {
	return &PaymentGrpcClient{
		client: paymentpb.NewPaymentServiceClient(conn),
	}
}

func (c *PaymentGrpcClient) ProcessPayment(ctx context.Context, orderID, buyerID uint, amount float64) (string, error) {
	return "", nil
}
