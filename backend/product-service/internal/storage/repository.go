package storage

import (
	"context"
	"product-service/internal/entities"
)

type ProductRepostitory interface {
	CreateProduct(ctx context.Context, pr *entities.Product) error
	GetProducts(ctx context.Context) ([]*entities.Product, error)
	GetProductById(ctx context.Context, id int64) (*entities.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
	UpdateProduct(ctx context.Context, id int64, updateFn func(*entities.Product) (bool, error)) error
}

type Repository struct {
	ProductRepo ProductRepostitory
}
