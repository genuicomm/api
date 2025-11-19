package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response adalah struktur dasar untuk semua response API
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Pagination adalah struktur untuk informasi pagination
type Pagination struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	From        int `json:"from"`
	To          int `json:"to"`
}

// Success mengirimkan response sukses
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Created mengirimkan response sukses untuk pembuatan resource baru
func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Error mengirimkan response error
func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Status:  "error",
		Message: message,
		Code:    statusCode,
	})
}

// ValidationError mengirimkan response error validasi
func ValidationError(c *gin.Context, message string, errors interface{}) {
	c.JSON(http.StatusUnprocessableEntity, Response{
		Status:  "error",
		Message: message,
		Code:    http.StatusUnprocessableEntity,
		Errors:  errors,
	})
}

// NotFound mengirimkan response not found
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Status:  "error",
		Message: message,
		Code:    http.StatusNotFound,
	})
}

// Unauthorized mengirimkan response unauthorized
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Response{
		Status:  "error",
		Message: message,
		Code:    http.StatusUnauthorized,
	})
}

// Forbidden mengirimkan response forbidden
func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, Response{
		Status:  "error",
		Message: message,
		Code:    http.StatusForbidden,
	})
}

// ServerError mengirimkan response server error
func ServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{
		Status:  "error",
		Message: message,
		Code:    http.StatusInternalServerError,
	})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Response{
		Status:  "error",
		Message: message,
		Code:    http.StatusBadRequest,
	})
}
