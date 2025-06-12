package grpcclient

import (
	"context"
	"fmt"
	"order-service/internal/adapters/grpc/gen/go/paymentpb"
	"order-service/internal/core/ports"

	"google.golang.org/grpc"
)

type PaymentGRPCClient struct {
	client paymentpb.PaymentServiceClient
}

func NewPaymentGRPCClient(conn *grpc.ClientConn) ports.PaymentService {
	return &PaymentGRPCClient{
		client: paymentpb.NewPaymentServiceClient(conn),
	}
}

func (c *PaymentGRPCClient) ProcessPayment(ctx context.Context, orderID, buyerID uint, amount float64) (string, error) {
	req := &paymentpb.PaymentRequest{
		BuyerID:   int32(buyerID),
		ProductID: int32(orderID),
		Amount:    amount,
	}
	res, err := c.client.ProcessPayment(ctx, req)
	if err != nil {
		return "", fmt.Errorf("grpc client error: %w", err)
	}

	return res.GetStatus(), nil
}
