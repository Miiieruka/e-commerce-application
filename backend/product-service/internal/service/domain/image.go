package domain

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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
	const op = "service.domain.image"
	file, err := hdr.Open()
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer file.Close()

	filename := fmt.Sprintf("event_%d%s", time.Now().UnixNano(), filepath.Ext(hdr.Filename))

	resp, err := imgSrv.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:     filename,
		Folder:       "products",
		ResourceType: "image",
	})

	fmt.Printf("%s and %s: resp\n", resp.URL, resp.SecureURL)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.SecureURL, nil
}
