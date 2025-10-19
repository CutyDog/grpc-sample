package model

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UID       string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
