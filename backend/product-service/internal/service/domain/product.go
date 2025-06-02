package domain

import (
	"context"
	"product-service/internal/entities"
	"product-service/internal/storage"
)

type ProductService struct {
	repo *storage.Repository
}

func NewProductService(repo *storage.Repository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, pr *entities.Product) error {
	return s.repo.ProductRepo.CreateProduct(ctx, pr)
}
func (s *ProductService) GetProducts(ctx context.Context) ([]*entities.Product, error) {
	return s.repo.ProductRepo.GetProducts(ctx)
}
func (s *ProductService) GetProductById(ctx context.Context, id int64) (*entities.Product, error) {
	return s.repo.ProductRepo.GetProductById(ctx, id)
}
func (s *ProductService) DeleteProduct(ctx context.Context, id int64) error {
	return s.repo.ProductRepo.DeleteProduct(ctx, id)
}
func (s *ProductService) UpdateProduct(ctx context.Context, id int64, pr *entities.ProductUpdateRequest) error {
	prod := &entities.Product{
		Name:        pr.Name,
		Description: pr.Description,
		Price:       pr.Price,
	}

	return s.repo.ProductRepo.UpdateProduct(ctx, id, prod)
}
