package repository

import (
	"BwiBOW123/backend-miniproject/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *UserRepository) LoginUser(username, password string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
