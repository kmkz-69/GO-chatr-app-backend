package models

import (

	"github.com/google/uuid"
	"gorm.io/gorm"

)

// AdminAttributes represents schema for admin table
type Admin struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Username  string   `gorm:"type:varchar(100);not null;unique"`
	Password  string   `gorm:"type:varchar(100);not null"`
	AccessToken string   `gorm:"type:varchar(100);not null"`
}

// CreateAdmin creates a new admin
