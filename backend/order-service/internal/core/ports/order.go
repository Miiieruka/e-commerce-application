package ports

import (
	"context"
	"order-service/internal/core/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
	GetByID(ctx context.Context, id uint) (*domain.Order, error)
	GetHistoryByBuyerID(ctx context.Context, buyerID uint) ([]domain.Order, error)
	UpdateStatus(ctx context.Context, id uint, status string) error
}

type OrderService interface {
	CreateOrder(ctx context.Context, buyerdId uint, orders []domain.OrderItem) (*domain.Order, error)
	GetOrderHistory(ctx context.Context, buyerID uint) ([]domain.Order, error)
}
