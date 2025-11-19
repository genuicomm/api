package user

import (
	"github.com/google/uuid"
	"github.com/genuicomm/api/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository membuat instance baru dari UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// GetUsers mengambil daftar user dengan filter dan pagination
func (r *userRepository) GetUsers(page, limit int, search, sort, order string) ([]domain.User, error) {
	var users []domain.User
	query := r.db.Model(&domain.User{})

	if search != "" {
		query = query.Where("username LIKE ? OR email LIKE ? OR full_name LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if sort != "" {
		query = query.Order(sort + " " + order)
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

// GetUserByID mengambil user berdasarkan ID
func (r *userRepository) GetUserByID(id uuid.UUID) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// GetUserByEmail mengambil user berdasarkan email
func (r *userRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// CreateUser membuat user baru
func (r *userRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

// UpdateUser memperbarui user yang ada
func (r *userRepository) UpdateUser(user *domain.User) error {
	return r.db.Save(user).Error
}

// DeleteUser menghapus user
func (r *userRepository) DeleteUser(id uuid.UUID) error {
	return r.db.Delete(&domain.User{}, "id = ?", id).Error
}
