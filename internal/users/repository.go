package users

import (
	"event-app/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]models.User, error)
	GetByID(id string) (*models.User, error)
	GetByEmail(id string) (*models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id string) (bool, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *repository) Create(user models.User) (models.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *repository) GetByID(id string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *repository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *repository) Update(user models.User) (models.User, error) {
	err := r.DB.Save(&user).Error
	return user, err
}

func (r *repository) Delete(id string) (bool, error) {
	err := r.DB.Delete(&models.User{}, "id = ?", id).Error
	return err == nil, err
}
