package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	const op = "repository.OrderRepository.Create"
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	query := `INSERT INTO orders (buyer_id, total_price, status, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err = tx.QueryRowContext(ctx, query, order.BuyerID, order.TotalPrice, order.Status, order.CreatedAt).Scan(&order.ID)
	println(order.ID)
	if err != nil {
		return fmt.Errorf("%s(insert order): %w", op, err)
	}
	for _, item := range order.Items {
		itemQuery := `INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)`
		_, err = tx.ExecContext(ctx, itemQuery, order.ID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return fmt.Errorf("%s(insert order_item): %w", op, err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (r *OrderRepository) GetByID(ctx context.Context, id uint) (*domain.Order, error) {
	const op = "repository.OrderRepository.GetByID"
	query := `SELECT id, buyer_id, total_price, status, created_at FROM orders WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	order := &domain.Order{}
	err := row.Scan(&order.ID, &order.BuyerID, &order.TotalPrice, &order.Status, &order.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: order not found", op)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	itemQuery := `SELECT product_id, quantity, price FROM order_items WHERE order_id = $1`
	rows, err := r.db.QueryContext(ctx, itemQuery, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		item := domain.OrderItem{}
		err = rows.Scan(&item.ProductID, &item.Quantity, &item.Price)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		order.Items = append(order.Items, item)
	}

	return order, nil
}
func (r *OrderRepository) GetHistoryByBuyerID(ctx context.Context, buyerID uint) ([]domain.Order, error) {
	const op = "repository.OrderRepository.GetHistoryByBuyerID"
	query := `SELECT id, buyer_id, total_price, status, created_at FROM orders WHERE buyer_id = $1 ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query, buyerID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		order := domain.Order{}
		err = rows.Scan(&order.ID, &order.BuyerID, &order.TotalPrice, &order.Status, &order.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		itemQuery := `SELECT product_id, quantity, price FROM order_items WHERE order_id = $1`
		itemRows, err := r.db.QueryContext(ctx, itemQuery, order.ID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		defer itemRows.Close()

		for itemRows.Next() {
			item := domain.OrderItem{}
			err = itemRows.Scan(&item.ProductID, &item.Quantity, &item.Price)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", op, err)
			}
			order.Items = append(order.Items, item)
		}

		orders = append(orders, order)
	}

	return orders, nil
}
func (r *OrderRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	const op = "repository.OrderRepository.UpdateStatus"
	query := `UPDATE orders SET status = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, status, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
