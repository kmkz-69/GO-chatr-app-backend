package repositories

import (
	"github.com/kmkz-69/chatr.fr.backend/src/models"
	"gorm.io/gorm"
)

// PrivacyRepository is the repository for privacy
type PrivacyRepository struct {
	db *gorm.DB
}

// CreatePrivacy creates a new privacy
func (pr *PrivacyRepository) CreatePrivacy(privacy *models.Privacy) error {
	return pr.db.Create(privacy).Error
}

// UpdatePrivacy updates a privacy
func (pr *PrivacyRepository) UpdatePrivacy(privacy *models.Privacy) error {
	return pr.db.Save(privacy).Error
}