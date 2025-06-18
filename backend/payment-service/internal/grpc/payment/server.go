package paymentgrpc

import (
	"context"
	"google.golang.org/grpc"
	"payment-service/internal/grpc/gen/go/paymentpb"
)

type serverAPI struct {
	paymentpb.UnimplementedPaymentServiceServer
}

func Register(gRPC *grpc.Server) {
	paymentpb.RegisterPaymentServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) ProcessPayment(
	ctx context.Context,
	in *paymentpb.PaymentRequest,
) (*paymentpb.PaymentResponse, error) {
	return &paymentpb.PaymentResponse{
		Status: "success",
	}, nil
}
