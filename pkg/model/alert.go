package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Alert struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	ReportedAt time.Time `json:"date"`
}

func (a *Alert) SaveAlert(db *gorm.DB) (*Alert, error) {
	err := db.Create(&a).Error
	if err != nil {
		return &Alert{}, err
	}

	return a, nil
}

func (a *Alert) GetAlert(db *gorm.DB, id uuid.UUID) (*Alert, error) {
	err := db.Model(Alert{}).Where("id = ?", id).Take(&a).Error
	if err != nil {
		return &Alert{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &Alert{}, errors.New("alert not found")
	}
	return a, nil
}

func (a *Alert) GetAllAlerts(db *gorm.DB) (*[]Alert, error) {
	var alerts []Alert
	err := db.Model(&Alert{}).Limit(100).Find(&alerts).Error
	if err != nil {
		return &[]Alert{}, err
	}

	return &alerts, err
}
