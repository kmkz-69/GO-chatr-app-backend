package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


// HostAttributes represents for hosts table
type Host struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Name      string   `gorm:"type:varchar(100);not null;unique"`
	Username string   `gorm:"type:varchar(100);not null"`
	Password string   `gorm:"type:varchar(100);not null"`
	AccessToken string   `gorm:"type:varchar(100);not null"`
	LastMeetingAt time.Time
}