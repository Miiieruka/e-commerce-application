package entities

import "time"

type Product struct {
	ID          int64
	SellerID    int64
	Name        string
	Description string
	Price       int64
	ImgUrl      string
	CreatedAt   time.Time
}
