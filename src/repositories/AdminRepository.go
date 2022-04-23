package repositories

import (
	"github.com/kmkz-69/chatr.fr.backend/src/models"
	"gorm.io/gorm"
)

// AdminRepository is the repository for admin
type AdminRepository struct {
	db *gorm.DB
}

// NewAdminRepository returns a new instance of AdminRepository
func Admin(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

// CreateAdmin creates a new admin
func (ar *AdminRepository) CreateAdmin(admin *models.Admin) error {
	return ar.db.Create(admin).Error
}

// GetAllAdmins returns all admins
func (ar *AdminRepository) GetAllAdmins() ([]*models.Admin, error) {
	admins := []*models.Admin{}
	err := ar.db.Find(&admins).Error
	return admins, err
}

// FindAdminByUsername finds an admin by username
func (ar *AdminRepository) FindAdminByUsername(username string) (*models.Admin, error) {
	admin := &models.Admin{}
	err := ar.db.Where("username = ?", username).First(admin).Error
	return admin, err
}

// FindAdminByAccessToken finds an admin by access token
func (ar *AdminRepository) FindAdminByAccessToken(accessToken string) (*models.Admin, error) {
	admin := &models.Admin{}
	err := ar.db.Where("access_token = ?", accessToken).First(admin).Error
	return admin, err
}