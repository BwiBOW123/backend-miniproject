package repository

import (
	"BwiBOW123/backend-miniproject/internal/domain"
	"crypto/sha256"
	"encoding/hex"

	"gorm.io/gorm"
)

func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *domain.Member) error {
	user.Password = hashPassword(user.Password)
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.Member, error) {
	var user domain.Member
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
func (r *UserRepository) GetUserByEmail(email string) (*domain.Member, error) {
	var user domain.Member
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) Login(username, password string) (*domain.Member, error) {
	var user domain.Member
	err := r.db.Where("username = ? AND password = SHA2(?, 256)", username, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
