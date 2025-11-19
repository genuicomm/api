package user

import (
	"github.com/google/uuid"
	"github.com/genuicomm/api/domain"
)

// UserUsecase mendefinisikan kontrak untuk operasi usecase user
type UserUsecase interface {
	GetUsers(page, limit int, search, sort, order string) ([]domain.User, error)
	GetUserByID(id uuid.UUID) (domain.User, error)
	CreateUser(user *domain.User) error
	UpdateUser(user *domain.User) error
	DeleteUser(id uuid.UUID) error
}

// UserRepository mendefinisikan kontrak untuk operasi repository user
type UserRepository interface {
	GetUsers(page, limit int, search, sort, order string) ([]domain.User, error)
	GetUserByID(id uuid.UUID) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	CreateUser(user *domain.User) error
	UpdateUser(user *domain.User) error
	DeleteUser(id uuid.UUID) error
}
