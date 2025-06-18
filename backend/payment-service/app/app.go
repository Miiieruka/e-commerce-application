package app

import (
	"log/slog"
	grpcapp "payment-service/app/grpc"
)

type GRPCServer struct {
	GRPCServer *grpcapp.App
}

func NewGRPCServer(log *slog.Logger, grpcPort int) *GRPCServer {
	grpcApp := grpcapp.New(log, grpcPort)

	return &GRPCServer{
		GRPCServer: grpcApp,
	}
}
