package users

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]User, error)
	GetByID(id string) (*User, error)
	Create(user User) (bool, error)
	Update(user User) (bool, error)
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

func (r *repository) Create(user User) (bool, error) {
	err := r.DB.Create(&user).Error
	return err == nil, err
}

func (r *repository) GetByID(id string) (*User, error) {
	var user User
	err := r.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *repository) Update(user User) (bool, error) {
	err := r.DB.Save(&user).Error
	return err == nil, err
}

func (r *repository) Delete(id string) (bool, error) {
	err := r.DB.Delete(&User{}, "id = ?", id).Error
	return err == nil, err
}
