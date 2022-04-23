package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PrivacyAttributes schema for privacy table
type Privacy struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Body      string   `gorm:"type:varchar(100);not null;unique"`
}
