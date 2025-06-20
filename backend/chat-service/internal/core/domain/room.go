package domain

import (
	"time"
)

type Room struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BuyerID     int
	SellerID    int
}
