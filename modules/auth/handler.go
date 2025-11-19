package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase AuthUsecase
}

// NewAuthHandler membuat instance baru dari AuthHandler
func NewAuthHandler(authUsecase AuthUsecase, router *gin.Engine) *AuthHandler {
	authHandler := &AuthHandler{
		authUsecase: authUsecase,
	}

	router.POST("/auth/login", authHandler.Login)

	return authHandler
}

// LoginRequest merepresentasikan request body untuk login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse merepresentasikan response untuk login
type LoginResponse struct {
	Token string `json:"token"`
}

// Login menangani request login
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.authUsecase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, response)
}
