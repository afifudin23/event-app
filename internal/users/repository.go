package users

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]User, error)
	GetByID(id string) (*User, error)
	GetByEmail(id string) (*User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(id string) (bool, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) GetAll() ([]User, error) {
	var users []User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *repository) Create(user User) (User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *repository) GetByID(id string) (*User, error) {
	var user User
	err := r.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *repository) GetByEmail(email string) (*User, error) {
	var user User
	err := r.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *repository) Update(user User) (User, error) {
	err := r.DB.Save(&user).Error
	return user, err
}

func (r *repository) Delete(id string) (bool, error) {
	err := r.DB.Delete(&User{}, "id = ?", id).Error
	return err == nil, err
}
