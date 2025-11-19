package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/genuicomm/api/domain"
)

type UserHandler struct {
	userUsecase UserUsecase
}

// NewUserHandler membuat instance baru dari UserHandler
func NewUserHandler(userUsecase UserUsecase, router *gin.RouterGroup) {
	userHandler := &UserHandler{
		userUsecase: userUsecase,
	}

	router.GET("user", userHandler.GetUsers)
	router.GET("user/:id", userHandler.GetUserByID)
	router.POST("user", userHandler.CreateUser)
	router.PUT("user/:id", userHandler.UpdateUser)
	router.DELETE("user/:id", userHandler.DeleteUser)
}

// GetUsers menangani request untuk mendapatkan daftar user
func (h *UserHandler) GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")
	sort := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "DESC")

	users, err := h.userUsecase.GetUsers(page, limit, search, sort, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID menangani request untuk mendapatkan user berdasarkan ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userUsecase.GetUserByID(uuid.MustParse(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUserRequest merepresentasikan request body untuk membuat user
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Role     string `json:"role"`
}

// CreateUser menangani request untuk membuat user baru
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &domain.User{
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
	}

	if err := h.userUsecase.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUserRequest merepresentasikan request body untuk memperbarui user
type UpdateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
	FullName string `json:"full_name" binding:"required"`
	Role     string `json:"role"`
}

// UpdateUser menangani request untuk memperbarui user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &domain.User{
		ID:       uuid.MustParse(id),
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
	}

	if err := h.userUsecase.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser menangani request untuk menghapus user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.userUsecase.DeleteUser(uuid.MustParse(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
