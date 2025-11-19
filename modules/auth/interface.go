package auth

import (
	"gitlab.com/genuicomm/api/domain"
)

// AuthRepository mendefinisikan kontrak untuk operasi repository auth
type AuthRepository interface {
	ValidateCredentials(username, password string) (*domain.User, error)
	GenerateToken(user *domain.User) (string, error)
	ValidateToken(token string) (*domain.User, error)
	RevokeToken(token string) error
}

// AuthUsecase mendefinisikan kontrak untuk operasi usecase auth
type AuthUsecase interface {
	Login(email, password string) (domain.LoginResponse, error)
	GenerateToken(userID string, email string) (string, error)
	GenerateRefreshToken(userID string, email string) (string, error)
}
