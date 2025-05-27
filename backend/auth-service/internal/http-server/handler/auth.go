package handler

import (
	"auth-service/internal/http-server/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) Register(c *gin.Context) {
	var req service.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.srv.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "User created succesfully!",
		"user_id": user.ID,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.srv.Login(req)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})

}

func (h *AuthHandler) UserInfo(c *gin.Context) {
	userID := c.GetString("user_id")
	role := c.GetString("role")

	c.IndentedJSON(http.StatusOK, gin.H{
		"user_id": userID,
		"role":    role,
	})
}

func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	url := service.GoogleOAuthConfig.AuthCodeURL("state-token")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) GoogleCallBack(c *gin.Context) {
	code := c.Query("code")
	token, err := service.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "token exchange failed"})
		return
	}
	jwtToken, err := h.srv.GoogleOAuthLogin(token)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": jwtToken})

}
