package users

import (
	"event-app/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll(loadRoles bool) ([]models.User, error)
	GetByID(id string, loadEvents bool, loadParticipations bool) (models.User, error)
	GetByEmail(email string, loadRoles bool) (models.User, error)
	GetRolesByUserID(userID string) ([]models.Role, error)
	GetRolesByName(roleName []string) ([]models.Role, error)
	GetRoleByName(roleName string) (models.Role, error)
	AssignRoleToUser(userRole models.UserRole) (bool, error)
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

func (r *repository) GetAll(loadRoles bool) ([]models.User, error) {
	var users []models.User
	query := r.DB

	if loadRoles {
		query = query.Preload("Roles")
	}
	err := query.Find(&users).Error
	return users, err
}

func (r *repository) Create(user models.User) (models.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *repository) GetByID(id string, loadEvents bool, loadParticipations bool) (models.User, error) {
	var user models.User
	query := r.DB

	if loadEvents {
		query = query.Preload("Events")
	}

	if loadParticipations {
		query = query.Preload("Participations.Event")
	}

	err := query.First(&user, "id = ?", id).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *repository) GetByEmail(email string, loadRoles bool) (models.User, error) {
	var user models.User
	query := r.DB

	if loadRoles {
		query = query.Preload("Roles")
	}
	err := query.First(&user, "email = ?", email).Error
	if err != nil {
		return models.User{}, err
	}
	return user, err
}

func (r *repository) GetRolesByUserID(userID string) ([]models.Role, error) {
	var roles []models.Role

	err := r.DB.
		Joins("JOIN user_roles ur ON ur.role_id = roles.id").
		Where("ur.user_id = ?", userID).
		Find(&roles).Error

	return roles, err
}

func (r *repository) GetRolesByName(roleName []string) ([]models.Role, error) {
	var roles []models.Role
	err := r.DB.Where("name IN ?", roleName).Find(&roles).Error
	return roles, err
}
func (r *repository) GetRoleByName(roleName string) (models.Role, error) {
	var role models.Role
	err := r.DB.First(&role, "name = ?", roleName).Error
	return role, err
}

func (r *repository) AssignRoleToUser(userRole models.UserRole) (bool, error) {
	err := r.DB.Create(&userRole).Error
	return err == nil, err
}

func (r *repository) Update(user models.User) (models.User, error) {
	err := r.DB.Save(&user).Error
	return user, err
}

func (r *repository) Delete(id string) (bool, error) {
	err := r.DB.Delete(&models.User{}, "id = ?", id).Error
	return err == nil, err
}
