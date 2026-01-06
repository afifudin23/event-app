package seeders

import (
	"event-app/internal/models"
	"log"
	"strings"

	"gorm.io/gorm"
)

type RolePermissionSeeder struct{}

func NewRolePermissionSeeder() Seeder {
	return &RolePermissionSeeder{}
}

var RolePermissions = map[string][]string{
	"superadmin": {
		"*",
	},
	"admin": {
		"events.*",
	},
	"user": {
		"events.read",
	},
}

func (s *RolePermissionSeeder) Run(db *gorm.DB) {
	log.Println("Running Role Permission Seeder...")

	for roleName, perms := range RolePermissions {
		// GET ROLE
		var role models.Role
		if err := db.First(&role, "name = ?", roleName).Error; err != nil {
			log.Fatalf("Role %s not found, run RoleSeeder first", roleName)
		}

		// GET PERMISSION
		var permissions []models.Permission
		permissionMap := make(map[string]models.Permission)

		if len(perms) == 1 && perms[0] == "*" {
			if err := db.Find(&permissions).Error; err != nil {
				log.Fatal(err)
			}
		} else {
			for _, perm := range perms {

				// WILDCARD users.*
				if base, ok := strings.CutSuffix(perm, ".*"); ok {
					prefix := base + "."
					var wildcardPerms []models.Permission
					if err := db.Where("name LIKE ?", prefix+"%").
						Find(&wildcardPerms).Error; err != nil {
						log.Fatal(err)
					}

					for _, p := range wildcardPerms {
						permissionMap[p.Name] = p
					}
					continue
				}

				// NORMAL permission
				var p models.Permission
				if err := db.First(&p, "name = ?", perm).Error; err != nil {
					log.Fatalf("Permission %s not found", perm)
				}
				permissionMap[p.Name] = p
			}

			for _, p := range permissionMap {
				permissions = append(permissions, p)
			}
		}

		// ASSIGN PERMISIONS TO ROLE
		if err := db.Model(&role).Association("Permisions").Replace(permissions); err != nil {
			log.Fatalf("Failed to assign permissions to role %s: %v", roleName, err)
		}
	}

	log.Println("Role Permission seeder completed")
}
