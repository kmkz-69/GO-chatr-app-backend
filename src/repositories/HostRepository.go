package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/kmkz-69/chatr.fr.backend/src/models"
	"gorm.io/gorm"
)

// HostRepository is the repository for hosts
type HostRepository struct {
	db *gorm.DB
}

// NewHostRepository returns a new instance of HostRepository
func NewHostRepository(db *gorm.DB) *HostRepository {
	return &HostRepository{db: db}
}

// CreateHostByUsernameAndPassword creates a new host by username and password 
func (hr *HostRepository) CreateHostByUsernameAndPassword(username, password string) (*models.Host, error) {
	host := &models.Host{
		Name: username,
		Username: username,
		Password: password,
	}
	err := hr.db.Create(host).Error
	return host, err
}

// FindHostByUsername finds a host by username 
func (hr *HostRepository) FindHostByUsername(username string) (*models.Host, error) {
	host := &models.Host{}
	err := hr.db.Where("username = ?", username).First(host).Error	
	return host, err
}

// FindHostByAccessToken finds a host by access token
func (hr *HostRepository) FindHostByAccessToken(accessToken string) (*models.Host, error) {
	host := &models.Host{}
	err := hr.db.Where("access_token = ?", accessToken).First(host).Error
	return host, err
}

// HostListAll returns all hosts
func (hr *HostRepository) HostListAll() ([]*models.Host, error) {
	hosts := []*models.Host{}
	err := hr.db.Find(&hosts).Error
	return hosts, err
}

// UpdateHostLastMeeting  by id 
func (hr *HostRepository) UpdateHostLastMeeting(id uuid.UUID) error {
	host := &models.Host{}
	err := hr.db.Where("id = ?", id).First(host).Error
	if err != nil {
		return err
	}
	host.LastMeetingAt = time.Now()
	return hr.db.Save(host).Error
}	

// DeleteHostByID deletes a host by id
func (hr *HostRepository) DeleteHostByID(id uuid.UUID) error {
	host := &models.Host{}
	err := hr.db.Where("id = ?", id).First(host).Error
	if err != nil {
		return err
	}
	return hr.db.Delete(host).Error
}