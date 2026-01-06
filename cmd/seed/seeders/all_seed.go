package seeders

import (
	"gorm.io/gorm"
)

type Seeder interface {
	Run(db *gorm.DB)
}

func RunAllSeeder(db *gorm.DB) {
	seeders := []Seeder{
		NewPermisionSeeder(),
		NewRoleSeeder(),
		NewRolePermissionSeeder(),
		NewUserSeeder(),
	}

	for _, seeder := range seeders {
		seeder.Run(db)
	}
}
