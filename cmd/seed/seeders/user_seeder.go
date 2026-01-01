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

func (s *UserSeeder) Run(db *gorm.DB) {
	log.Println("Running User Seeder...")
	passwordHashed, err := security.HashPassword("admin")
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	user := models.User{
		Fullname: "Admin",
		Email:    "admin@example.com",
		Password: passwordHashed,
	}

	var existingUser models.User
	result := db.Where("email = ?", user.Email).First(&existingUser)

	if result.Error == gorm.ErrRecordNotFound {
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Failed to create user admin: %v", err)
		}
		log.Println("Created user admin successfully")
	} else if result.Error != nil {
		log.Fatalf("Failed to check if user admin exists: %v", result.Error.Error())
	} else {
		log.Println("User admin already exists")
	}
}
