package repository

import (
	"todolist-backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByUsername(username string, user *model.User) error {
	return r.DB.Where("username = ?", username).First(user).Error
}

func (r *UserRepository) FindByID(id uint, user *model.User) error {
	return r.DB.Where("id = ?", id).First(user).Error
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User

	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
