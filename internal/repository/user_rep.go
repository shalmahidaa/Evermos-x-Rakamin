package repository

import (
	"evermos/config"
	"evermos/internal/domain"
)

type userRepository struct{}

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}

func (r *userRepository) Save(user *domain.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := config.DB.Preload("Store").Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByPhone(phone string) (*domain.User, error) {
	var user domain.User
	err := config.DB.Where("phone = ?", phone).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := config.DB.Preload("Store").First(&user, id).Error
	return &user, err
}

func (r *userRepository) Update(userID uint, updated *domain.User) error {
	return config.DB.Model(&domain.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"name":  updated.Name,
		"email": updated.Email,
		"phone": updated.Phone,
	}).Error
}