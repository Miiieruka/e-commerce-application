package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"order-service/internal/core/domain"
	"order-service/internal/core/ports"
	"time"
)

type OrderHandler struct {
	srv ports.OrderService
}

func NewOrderHandler(srv ports.OrderService) *OrderHandler {
	return &OrderHandler{
		srv: srv,
	}
}

type CreateOrderRequest struct {
	Items []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order CreateOrderRequest
	buyerID := c.GetInt64("user_id")
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var orderItems []domain.OrderItem
	for _, item := range order.Items {
		orderItems = append(orderItems, domain.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}
	responseOrder, err := h.srv.CreateOrder(ctx, uint(buyerID), orderItems)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(201, gin.H{"order_id": responseOrder.ID})
}

func (h *OrderHandler) GetOrderHistory(c *gin.Context) {
	buyerID := c.GetInt64("user_id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	orders, err := h.srv.GetOrderHistory(ctx, uint(buyerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order history"})
		return
	}

	c.JSON(200, gin.H{"orders": orders})
}
