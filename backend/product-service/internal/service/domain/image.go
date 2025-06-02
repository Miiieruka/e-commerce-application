package domain

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
)

type ImageService struct {
	cld *cloudinary.Cloudinary
}

func NewImageService(cld *cloudinary.Cloudinary) *ImageService {
	return &ImageService{
		cld: cld,
	}
}

func (imgSrv *ImageService) UploadImage(ctx context.Context, hdr *multipart.FileHeader) (string, error) {
	return "", nil
}
