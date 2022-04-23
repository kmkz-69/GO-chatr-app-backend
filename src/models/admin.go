package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

)

// AdminAttributes represents schema for admin table
type AdminAttributes struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Username  string   `gorm:"type:varchar(100);not null;unique"`
	Password  string   `gorm:"type:varchar(100);not null"`
	AccessToken string   `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time	
	UpdatedAt time.Time
}

// CreateAdmin creates a new admin
