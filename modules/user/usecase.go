package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/genuicomm/api/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repo UserRepository
}

// NewUserUsecase membuat instance baru dari UserUsecase
func NewUserUsecase(repo UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

// GetUsers mengambil daftar user dengan filter dan pagination
func (u *userUsecase) GetUsers(page, limit int, search, sort, order string) ([]domain.User, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if sort == "" {
		sort = "created_at"
	}
	if order == "" {
		order = "DESC"
	}

	return u.repo.GetUsers(page, limit, search, sort, order)
}

// GetUserByID mengambil user berdasarkan ID
func (u *userUsecase) GetUserByID(id uuid.UUID) (domain.User, error) {
	if id == uuid.Nil {
		return domain.User{}, errors.New("invalid user ID")
	}

	return u.repo.GetUserByID(id)
}

// CreateUser membuat user baru
func (u *userUsecase) CreateUser(user *domain.User) error {
	// Validasi input
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password are required")
	}

	// Cek apakah email sudah ada
	existingUser, err := u.repo.GetUserByEmail(user.Email)
	if err == nil && existingUser.ID != uuid.Nil {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = hashedPassword

	// Set timestamp
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	return u.repo.CreateUser(user)
}

// UpdateUser memperbarui user yang ada
func (u *userUsecase) UpdateUser(user *domain.User) error {
	if user.ID == uuid.Nil {
		return errors.New("invalid user ID")
	}

	// Cek apakah user ada
	existingUser, err := u.repo.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	// Update timestamp
	user.UpdatedAt = time.Now()

	// Jika password diubah, hash password baru
	if user.Password != "" && user.Password != existingUser.Password {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	} else {
		user.Password = existingUser.Password
	}

	return u.repo.UpdateUser(user)
}

// DeleteUser menghapus user
func (u *userUsecase) DeleteUser(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}

	// Cek apakah user ada
	_, err := u.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	return u.repo.DeleteUser(id)
}
