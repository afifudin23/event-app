package seeders

import (
	"event-app/internal/models"
	"log"

	"gorm.io/gorm"
)

type RoleSeeder struct{}

func NewRoleSeeder() Seeder {
	return &RoleSeeder{}
}

var Roles = []string{
	"superadmin",
	"admin",
	"user",
}

func (s *RoleSeeder) Run(db *gorm.DB) {
	log.Println("Running Role Seeder...")

	for _, name := range Roles {
		role := models.Role{Name: name}

		if err := db.FirstOrCreate(&role, models.Role{Name: name}).Error; err != nil {
			log.Fatalf("Failed to create role name '%s': %v", name, err)
		}
	}

	log.Println("Role seeder completed")
}
