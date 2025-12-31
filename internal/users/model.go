package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Fullname  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);not null;unique"`
	Password  string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`
}
