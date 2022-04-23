package services

import "gorm.io/gorm"

// Validation is the validation service
type Validation struct {
	db *gorm.DB
}

// func for validation data handle error 
func (v *Validation) handleError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
