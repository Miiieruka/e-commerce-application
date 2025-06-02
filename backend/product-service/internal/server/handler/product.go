package handler

import (
	"context"
	"fmt"
	"net/http"
	"product-service/internal/entities"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SaveProduct(c *gin.Context) {
	if c.GetString("role") != "seller" {
		c.JSON(http.StatusForbidden, gin.H{"error": "only sellers can add products"})
		return
	}

	image, err := c.FormFile("img_url")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image is required"})
		return
	}
	fmt.Printf("%s", image.Filename)
	name := c.PostForm("name")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid price"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	url, err := h.services.ImageService.UploadImage(ctx, image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload image"})
		fmt.Printf("save product: %s\n", err.Error())
		return
	}

	prod := &entities.Product{
		Name:        name,
		Description: description,
		Price:       price,
		ImgUrl:      url,
		SellerID:    c.GetInt64("user_id"),
		CreatedAt:   time.Now(),
	}

	if err := h.services.ProductService.CreateProduct(ctx, prod); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save product"})
		fmt.Printf("save product: %s\n", err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "product added"})
}

func (h *Handler) GetAllProducts(c *gin.Context) {
	ctxWithDeadline, cancel := context.WithTimeout(c.Request.Context(), time.Second*2)
	defer cancel()
	products, err := h.services.ProductService.GetProducts(ctxWithDeadline)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "couldn't fetch products"})
		fmt.Printf("get all product: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, products)
}

func (h *Handler) GetProductById(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	ctxWithDeadline, cancel := context.WithTimeout(c.Request.Context(), time.Second*2)
	defer cancel()
	product, err := h.services.ProductService.GetProductById(ctxWithDeadline, id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "product not found"})
		fmt.Printf("get product: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	if c.GetString("role") != "seller" {
		c.IndentedJSON(http.StatusForbidden, gin.H{"error": "only sellers can update products"})
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	ctxWithDeadline, cancel := context.WithTimeout(c.Request.Context(), time.Second*2)
	defer cancel()
	var updatedProduct entities.ProductUpdateRequest
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	fmt.Printf("%s\n", updatedProduct.Description)

	err = h.services.ProductService.UpdateProduct(ctxWithDeadline, id, &updatedProduct)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "couldn't update product"})
		fmt.Printf("update product: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "product updated"})

}

func (h *Handler) DeleteProduct(c *gin.Context) {
	if c.GetString("role") != "seller" {
		c.JSON(http.StatusForbidden, gin.H{"error": "only sellers can delete products"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	ctxWithDeadline, cancel := context.WithTimeout(c.Request.Context(), time.Second*2)
	defer cancel()
	if err := h.services.ProductService.DeleteProduct(ctxWithDeadline, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		fmt.Printf("delete product: %s\n", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
