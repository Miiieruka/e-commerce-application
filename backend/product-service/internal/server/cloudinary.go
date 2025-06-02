package server

import (
	"log"
	"product-service/config"

	"github.com/cloudinary/cloudinary-go/v2"
)

func InitCloudinary(cloudyCfg config.CloudinaryConfig) *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(cloudyCfg.Name, cloudyCfg.ApiKey, cloudyCfg.ApiSecret)
	if err != nil {
		log.Fatalf("Init cloudy error: %s", err.Error())
	}
	return cld
}
