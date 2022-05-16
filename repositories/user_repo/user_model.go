package user_repo

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	//ID        uuid.UUID
	FullName  string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
