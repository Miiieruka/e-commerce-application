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
	Price       int64
	ImgUrl      string
	CreatedAt   time.Time
}

type ProductRequest struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Price       int64                 `json:"price"`
	Image       *multipart.FileHeader `json:"img_url"`
	SellerId    int64
}

type ProductUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}
