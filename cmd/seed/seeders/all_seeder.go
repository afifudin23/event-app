package seeders

import (
	"gorm.io/gorm"
)

type Seeder interface {
	Run(db *gorm.DB)
}

func RunAllSeeder(db *gorm.DB) {
	seeders := []Seeder{
		NewUserSeeder(),
	}

	for _, seeder := range seeders {
		seeder.Run(db)
	}
}
