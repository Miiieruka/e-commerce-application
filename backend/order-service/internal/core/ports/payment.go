package ports

import "context"

type PaymentService interface {
	ProcessPayment(ctx context.Context, orderID, buyerID uint, amount float64) (string, error)
}
