package service

import (
	"context"
	"fmt"
	"order-service/internal/core/domain"
	"order-service/internal/core/ports"
	"time"
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
	const op = "core.service.createorder"

	if len(items) == 0 {
		return nil, fmt.Errorf("%s: No items in order", op)
	}

	var total float64
	for _, item := range items {
		total += float64(item.Quantity) * item.Price
	}

	order := &domain.Order{
		BuyerID:    buyerID,
		Items:      items,
		TotalPrice: total,
		Status:     "pending",
		CreatedAt:  time.Now(),
	}

	err := s.repo.Create(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	status, err := s.payment.ProcessPayment(ctx, order.ID, buyerID, total)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = s.repo.UpdateStatus(ctx, order.ID, status)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	order.Status = status
	return order, nil
}

func (s *orderService) GetOrderHistory(ctx context.Context, buyerID uint) ([]domain.Order, error) {
	return s.repo.GetHistoryByBuyerID(ctx, buyerID)
}
