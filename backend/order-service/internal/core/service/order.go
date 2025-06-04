package service

import (
	"context"
	"order-service/internal/core/domain"
	"order-service/internal/core/ports"
)

type orderService struct {
	repo    ports.OrderRepository
	payment ports.PaymentService
}

func NewOrderService(repo ports.OrderRepository, payment ports.PaymentService) ports.OrderService {
	return &orderService{
		repo:    repo,
		payment: payment,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, buyerID uint, items []domain.OrderItem) (*domain.Order, error) {
	return nil, nil
}

func (s *orderService) GetOrderHistory(ctx context.Context, buyerID uint) ([]domain.Order, error) {
	return nil, nil
}
