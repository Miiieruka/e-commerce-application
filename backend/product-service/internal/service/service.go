package service

import (
	"context"
	"mime/multipart"
	"product-service/internal/entities"
	"product-service/internal/service/domain"
	"product-service/internal/storage"

	"github.com/cloudinary/cloudinary-go/v2"
)

type ProductService interface {
	CreateProduct(ctx context.Context, pr *entities.Product) error
	GetProducts(ctx context.Context) ([]*entities.Product, error)
	GetProductById(ctx context.Context, id int64) (*entities.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
	UpdateProduct(ctx context.Context, id int64, pr *entities.ProductUpdateRequest) error
}

type Imageservice interface {
	UploadImage(ctx context.Context, hdr *multipart.FileHeader) (string, error)
}

type Service struct {
	ProductService ProductService
	ImageService   Imageservice
}

func NewService(repo *storage.Repository, cld *cloudinary.Cloudinary) *Service {
	return &Service{
		ProductService: domain.NewProductService(repo),
		ImageService:   domain.NewImageService(cld),
	}
}
