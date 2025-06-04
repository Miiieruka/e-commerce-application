package domain

import "time"

type OrderItem struct {
	ID        uint
	ProductID uint
	Price     float64
	Quantity  uint
}

type Order struct {
	ID         uint
	BuyerID    uint
	Items      []OrderItem
	TotalPrice float64
	Status     string
	CreatedAt  time.Time
}
