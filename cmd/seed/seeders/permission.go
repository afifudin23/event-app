package seeders

import (
	"event-app/internal/models"
	"log"

	"gorm.io/gorm"
)

type PermisionSeeder struct{}

func NewPermisionSeeder() Seeder {
	return &PermisionSeeder{}
}

var Permissions = []string{
	"users.read",
	"users.create",
	"users.update",
	"users.delete",

	"roles.read",
	"roles.create",
	"roles.update",
	"roles.delete",
	"roles.assign_permission",

	"events.read",
	"events.create",
	"events.update",
	"events.delete",

	"event_participants.read",
	"event_participants.register",
	"event_participants.cancel",
}

func (s *PermisionSeeder) Run(db *gorm.DB) {
	log.Println("Running Permission Seeder...")
	for _, name := range Permissions {
		permission := models.Permission{Name: name}

		if err := db.FirstOrCreate(&permission, models.Permission{Name: name}).Error; err != nil {
			log.Fatalf("Failed to create permission name '%s': %v", name, err)
		}
	}
	log.Println("Permission seeder completed")
}
