package repository

import (
	"context"
	"database/sql"
	"order-service/internal/core/domain"
	"order-service/internal/core/ports"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) ports.OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Create(ctx context.Context, order *domain.Order) error {
	return nil
}
func (r *OrderRepository) GetByID(ctx context.Context, id uint) (*domain.Order, error) {
	return nil, nil
}
func (r *OrderRepository) GetHistoryByBuyerID(ctx context.Context, buyerID uint) ([]domain.Order, error) {
	return nil, nil
}
func (r *OrderRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	return nil
}
