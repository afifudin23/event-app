package seeders

import (
	"log"

	"event-app/internal/models"
	"event-app/pkg/security"

	"gorm.io/gorm"
)

type UserSeeder struct{}

func NewUserSeeder() Seeder {
	return &UserSeeder{}
}

type UserModel struct {
	Fullname string
	Email    string
	Password string
	Roles    []string
}

var Users = []UserModel{
	{
		Fullname: "Superadmin",
		Email:    "superadmin@example.com",
		Password: "superadmin",
		Roles:    []string{"superadmin"},
	},
	{
		Fullname: "Admin",
		Email:    "admin@example.com",
		Password: "admin",
		Roles:    []string{"admin"},
	},
	{
		Fullname: "User",
		Email:    "user@example.com",
		Password: "user",
		Roles:    []string{"user"},
	},
}

func (s *UserSeeder) Run(db *gorm.DB) {
	log.Println("Running User Seeder...")

	for _, user := range Users {
		s.CreateUser(db, user)
	}

	log.Println("Created user admin successfully")
}

func (s *UserSeeder) CreateUser(db *gorm.DB, u UserModel) {
	passwordHashed, err := security.HashPassword(u.Password)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	var user models.User
	if err := db.
		Where("email = ?", u.Email).
		Assign(models.User{
			Fullname: u.Fullname,
			Password: passwordHashed,
			Email:    u.Email,
		}).
		FirstOrCreate(&user).Error; err != nil {
		log.Fatalf("Failed to seed user email '%s': %v", u.Email, err)
	}

	var roles []models.Role
	if err := db.Find(&roles, "name IN ?", u.Roles).Error; err != nil {
		log.Fatalf("Failed to fetch roles for %s: %v", user.Email, err)
	}

	if len(roles) != len(u.Roles) {
		log.Fatalf("Warning: Not all roles found for %s. Found: %v, Requested: %v", u.Email, roles, u.Roles)
	}

	if err := db.Model(&user).Association("Roles").Replace(&roles); err != nil {
		log.Fatalf("Failed to assign roles to user %s: %v", user.Email, err)
	}
}
