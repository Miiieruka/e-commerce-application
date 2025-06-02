package entities

import (
	"mime/multipart"
	"time"
)

type Product struct {
	ID          int64
	SellerID    int64
	Name        string
	Description string
	Price       float64
	ImgUrl      string
	CreatedAt   time.Time
}

type ProductRequest struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Price       float64               `json:"price"`
	Image       *multipart.FileHeader `json:"img_url"`
}

type ProductUpdateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
