package models

import (
	"time"

	"gorm.io/gorm"
)

type AdminStatus string

const (
	IsActive   AdminStatus = "active"
	IsInactive AdminStatus = "inactive"
)

type Admin struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Firstname string         `gorm:"type:varchar(100);not null" json:"firstname"`
	Lastname  string         `gorm:"type:varchar(100);not null" json:"lastname"`
	Email     string         `gorm:"not null;uniqueIndex" json:"email"`
	Password  string         `gorm:"type:varchar(160);not null" json:"-"`
	Status    AdminStatus    `gorm:"type:varchar(20);not null;default:'inactive'" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"` // soft delete, hidden from API
}
